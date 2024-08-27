from fastapi import FastAPI, HTTPException
from celery import Celery
from pydantic import BaseModel
from datetime import datetime
import requests

app = FastAPI()

# Celery configuration
celery_app = Celery('tasks', broker='redis://localhost:6379/0')

class ReminderRequest(BaseModel):
    task_id: str
    title: str
    remind_at: datetime

@celery_app.task
def send_reminder(task_id: str, title: str):
    print(f"REMINDER: Task {task_id} - {title} is due soon!")
    
    # Send notification to Golang API
    notification_url = "http://localhost:8000/notifications"  # Updated to match your Golang API port
    notification_data = {
        "task_id": task_id,
        "title": title
    }
    
    try:
        response = requests.post(notification_url, json=notification_data)
        response.raise_for_status()
        print(f"Notification sent successfully for task {task_id}")
    except requests.exceptions.RequestException as e:
        print(f"Failed to send notification for task {task_id}: {e}")

@app.post("/schedule_reminder/")
async def schedule_reminder(reminder: ReminderRequest):
    try:
        # Schedule the reminder
        send_reminder.apply_async(args=[reminder.task_id, reminder.title], eta=reminder.remind_at)
        return {"message": "Reminder scheduled"}
    except Exception as e:
        print(f"Error scheduling reminder: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

@app.get("/")
async def root():
    return {"message": "Task Scheduler Service"}
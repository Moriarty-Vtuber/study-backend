# My Study Space (Go + Supabase + Render)

This is a migration of the Study Space bot to a robust Go backend hosted on Render with a Supabase database.

## Setup Instructions

### 1. Supabase (Database)
1. Create a free project at [Supabase](https://supabase.com).
2. Go to **Project Settings > Database** and copy the **Connection String (URI)**.
3. Go to the **SQL Editor** and run the following SQL to create the tables:

```sql
-- Users table
CREATE TABLE users (
    id TEXT PRIMARY KEY, -- YouTube Channel ID
    name TEXT NOT NULL,
    display_name TEXT,
    profile_image_url TEXT,
    total_study_time BIGINT DEFAULT 0, -- In seconds
    last_entered_at TIMESTAMP WITH TIME ZONE
);

-- Current Room State
CREATE TABLE room_state (
    user_id TEXT PRIMARY KEY REFERENCES users(id),
    seat_id INT, -- 1 to N
    entered_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    status TEXT DEFAULT 'studying' -- 'studying', 'break'
);
```

### 2. Google Cloud (YouTube API)
1. Create a project in Google Cloud Console.
2. Enable **YouTube Data API v3**.
3. Create an **API Key**.

### 3. Local Development
1. Install Go 1.21+.
2. Run `go mod tidy` to install dependencies.
3. Set environment variables:
   - `DATABASE_URL`: Your Supabase connection string.
   - `YOUTUBE_API_KEY`: Your Google Cloud API Key.
   - `VIDEO_ID`: The ID of your active live stream.
4. Run the server:
   ```bash
   go run cmd/server/main.go
   ```

### 4. Deployment (Render)
1. Push this repository to GitHub.
2. Create a new **Web Service** on Render.
3. Connect your repo.
4. Set Environment Variables in Render:
   - `DATABASE_URL`
   - `YOUTUBE_API_KEY`
   - `VIDEO_ID`
5. Deploy!

### 5. OBS Overlay
- Add a **Browser Source** in OBS.
- URL: `https://your-app-name.onrender.com/overlay/index.html`
- Width: 1920, Height: 1080.

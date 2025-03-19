package database

import "time"

// Reminders structure for the reminders in the database
type Reminder struct {
	ReminderID      int64     `db:"reminder_id"`
	ServerID        int64     `db:"server_id"`
	UserID          int64     `db:"user_id"`
	Message         string    `db:"message"`
	ChannelID       int64     `db:"channel_id"`
	Timestamp       time.Time `db:"timestamp"`
	Status          string    `db:"status"`
	SnoozeRemaining int       `db:"snooze_remaining"`
	CreatedAt       time.Time `db:"created_at"`
}

// BotLog is a log entry
type BotLog struct {
	LogID     int64     `db:"log_id"`
	ServerID  int64     `db:"server_id"`
	UserID    int64     `db:"user_id"`
	Timestamp time.Time `db:"timestamp"`
	LogType   string    `db:"log_type"`
	EventName string    `db:"event_name"`
	Message   string    `db:"message"`
	CreatedAt time.Time `db:"created_at"`
}

// UserConsent represents the permissions between users to be/get pinged in a server
type UserConsent struct {
	ServerID      int64     `db:"server_id"`
	UserID        int64     `db:"user_id"`
	AllowedUSerID int64     `db:"allowed_user_id"`
	CreatedAt     time.Time `db:"created_at"`
}

// ServerSettings stores per server settings
type ServerSettings struct {
	ServerID      int64 `db:"server_id"`
	UpdateChannel int64 `db:"update_channel"`
}

// SeverAllowedChannel represents which channels are allowed to message in
type ServerAllowedChannels struct {
	ServerID  int64 `db:"sever_id"`
	ChannelID int64 `db:"channel_id"`
}

type UserSettings struct {
	UserID                          int64  `db:"user_id"`
	Timezone                        string `db:"timezone"`
	DefaultDroppedReminderBehaviour string `db:"default_dropped_reminder_behaviour"`
}

CREATE TABLE IF NOT EXISTS event_participants(
    event_id SERIAL NOT NULL,
    participant_id SERIAL NOT NULL,
    checkin_at TIMESTAMP,
    PRIMARY KEY (event_id, participant_id),
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
    FOREIGN KEY (participant_id) REFERENCES participants(id) ON DELETE CASCADE
);

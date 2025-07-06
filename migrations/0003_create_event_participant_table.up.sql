CREATE TABLE event_participants (
    event_id INTEGER NOT NULL,
    participant_id INTEGER NOT NULL,
    checked_in BOOLEAN DEFAULT FALSE,
    checkin_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (event_id, participant_id),

    CONSTRAINT fk_event
        FOREIGN KEY (event_id)
        REFERENCES events (id)
        ON DELETE CASCADE,

    CONSTRAINT fk_participant
        FOREIGN KEY (participant_id)
        REFERENCES participants (id)
        ON DELETE CASCADE
);

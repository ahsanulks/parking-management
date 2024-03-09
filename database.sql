-- Create ParkingLot table
CREATE TABLE IF NOT EXISTS parking_lots (
    id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL, capacity INT NOT null, slot_left INT NOT null, manager_id INT NOT null
);

-- Create ParkingSlot table
CREATE TABLE IF NOT EXISTS parking_slots (
    id SERIAL PRIMARY KEY, parking_lot_id INT NOT NULL, "number" INT NOT NULL, available BOOLEAN NOT NULL DEFAULT TRUE, in_maintenance BOOLEAN NOT NULL DEFAULT FALSE, FOREIGN KEY (parking_lot_id) REFERENCES parking_lots (id)
);

-- Create Ticket table
CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY, parking_lot_id INT NOT NULL, parking_slot_id INT NOT NULL, code VARCHAR(255) NOT NULL, entry_time TIMESTAMP NOT NULL, exit_time TIMESTAMP, fee INT NULL, hours INT NULL, FOREIGN KEY (parking_lot_id) REFERENCES parking_lots (id), FOREIGN KEY (parking_slot_id) REFERENCES parking_slots (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tickets_code ON tickets (code);

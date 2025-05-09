-- Insert consumers
INSERT INTO consumers (id, nik, full_name, legal_name, place_of_birth, date_of_birth, salary, photo_ktp, photo_selfie)
VALUES
    ('11111111-1111-1111-1111-111111111111', '1234567890123456', 'Budi Santoso', 'Budi Santoso', 'Jakarta', '1990-01-10', 5000000, 'budi_ktp.jpg', 'budi_selfie.jpg'),
    ('22222222-2222-2222-2222-222222222222', '6543210987654321', 'Annisa Putri', 'Annisa Putri', 'Bandung', '1992-05-20', 8000000, 'annisa_ktp.jpg', 'annisa_selfie.jpg');

-- Insert limits for Budi
INSERT INTO limits (consumer_id, tenor, amount)
VALUES
    ('11111111-1111-1111-1111-111111111111', 1, 100000),
    ('11111111-1111-1111-1111-111111111111', 2, 200000),
    ('11111111-1111-1111-1111-111111111111', 3, 500000),
    ('11111111-1111-1111-1111-111111111111', 6, 700000);

-- Insert limits for Annisa
INSERT INTO limits (consumer_id, tenor, amount)
VALUES
    ('22222222-2222-2222-2222-222222222222', 1, 1000000),
    ('22222222-2222-2222-2222-222222222222', 2, 1200000),
    ('22222222-2222-2222-2222-222222222222', 3, 1500000),
    ('22222222-2222-2222-2222-222222222222', 6, 2000000);

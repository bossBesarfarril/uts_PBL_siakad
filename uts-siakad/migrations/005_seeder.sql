-- 1. SEEDER USERS
INSERT INTO users (email, password, role) VALUES ('admin@siakad.ac.id', '$2a$10$w24ZQ/mII5EioLtpHtjYI.Lq58FqCcssinCKPTWe7hI1mrhOe7S3y', 'admin');
INSERT INTO users (email, password, role) VALUES ('mhs1@siakad.ac.id', '$2a$10$zB7SlfC0uCZPznzpWLmTBeZkr7OPP.7oWHK1tDFo009/4.mJG1HE2', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs2@siakad.ac.id', '$2a$10$rloi90aEe6WuJ7a2au8ly.j2ze3PUwvDQG//CEJwhiwdZXHPBTz4a', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs3@siakad.ac.id', '$2a$10$FrfsKFJ8WvCh1gWtptQAHeJlVJhRHTkZ4RNzJAzY.q7sePC/xC6P.', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs4@siakad.ac.id', '$2a$10$3Xk7.0DgI6DsBPJ0O1GzxO0Z71hjHB/Br/sKgYPZmA9pZB7Ru4DlS', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs5@siakad.ac.id', '$2a$10$fcZPDehLGa8zYzXrdLFe7.cqnUgjC5PSXneNvtt7MtpukylILHUaW', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs6@siakad.ac.id', '$2a$10$aSMHyJ6ze0A.XXV9MLONUug55f2v6jOpY03yd1xEthe4xT/rQZL82', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs7@siakad.ac.id', '$2a$10$GGvO3Wv4GKEIlHNGwV/qL.kPmgmkwk9RkDgRJdDFVvuR2px4H4qFq', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs8@siakad.ac.id', '$2a$10$aEdS3L4D3phaPQ/tAtdIh.MKudk0ywyPFGyTUqBhkkiR4GU1ygqCS', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs9@siakad.ac.id', '$2a$10$1g1CJOydHl49amXc/k2zgO8QCOJis8y8u3OveXyYdf/DCuPVWz9gu', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs10@siakad.ac.id', '$2a$10$PZNlxDE7cXANFoTVFNMv3eY.VUe3TX.tjqCMqE.sdy/JFpznrFRwS', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs11@siakad.ac.id', '$2a$10$v3GyURCm8qSOx4A.Rck.uuuYE6JSDjGyYn/775DSGHLs6amn9Q58u', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs12@siakad.ac.id', '$2a$10$m5HRD0UbsBobXfeb/LpHfOX/M7LL.Ct23lUhSDemIamgI58RZGGv.', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs13@siakad.ac.id', '$2a$10$PSytf9lHcVrZ3Avm/uV4BuncvuyP0.zx3M0bTVMsWaPgv6Fg0Jwu.', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs14@siakad.ac.id', '$2a$10$DOHZGkHIXTiLSICdStIgH.St5h2WSop0CGiQGZaHJD9zjiVEVsQGm', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs15@siakad.ac.id', '$2a$10$lovOvgqWzWZj.Km6KnqZ9OSX4OslHIbmTF.P/hfx42sDcoMpZoalW', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs16@siakad.ac.id', '$2a$10$YnrFwz9k/Fu7G3Byne2i..1.eS.LP7Et2Ar/mMFsNRNeIdGDZfJq2', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs17@siakad.ac.id', '$2a$10$jotz1nIeA7NrOBAvssPnQeO17pvGLteONC3KnymseYvx3XxSc8UvG', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs18@siakad.ac.id', '$2a$10$HWIGSRI4lUIvAg94bhFrdOFGEckexeavnt6FXBPpYLhY4SlHmAiIe', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs19@siakad.ac.id', '$2a$10$L2YybM5pT20cq2iju1JFOuAM.BC8DIwx9DnSJkJotk.LZCH6w0i62', 'mahasiswa');
INSERT INTO users (email, password, role) VALUES ('mhs20@siakad.ac.id', '$2a$10$BJ7Xx8CvCSLPIITKm.rPFuPPCkf8i7WZEgAQInwhb9BbD5Jb1ebOe', 'mahasiswa');

-- 2. SEEDER STUDENTS
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (2, '123456789001', 'Mahasiswa 1', 'Sistem Informasi', 2023, 2.10);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (3, '123456789002', 'Mahasiswa 2', 'Informatika', 2023, 2.20);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (4, '123456789003', 'Mahasiswa 3', 'Sistem Informasi', 2023, 2.30);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (5, '123456789004', 'Mahasiswa 4', 'Informatika', 2023, 2.40);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (6, '123456789005', 'Mahasiswa 5', 'Sistem Informasi', 2023, 2.50);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (7, '123456789006', 'Mahasiswa 6', 'Informatika', 2023, 2.60);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (8, '123456789007', 'Mahasiswa 7', 'Sistem Informasi', 2023, 2.70);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (9, '123456789008', 'Mahasiswa 8', 'Informatika', 2023, 2.80);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (10, '123456789009', 'Mahasiswa 9', 'Sistem Informasi', 2023, 2.90);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (11, '123456789010', 'Mahasiswa 10', 'Informatika', 2023, 3.00);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (12, '123456789011', 'Mahasiswa 11', 'Sistem Informasi', 2024, 3.10);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (13, '123456789012', 'Mahasiswa 12', 'Informatika', 2024, 3.20);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (14, '123456789013', 'Mahasiswa 13', 'Sistem Informasi', 2024, 3.30);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (15, '123456789014', 'Mahasiswa 14', 'Informatika', 2024, 3.40);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (16, '123456789015', 'Mahasiswa 15', 'Sistem Informasi', 2024, 3.50);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (17, '123456789016', 'Mahasiswa 16', 'Informatika', 2024, 3.60);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (18, '123456789017', 'Mahasiswa 17', 'Sistem Informasi', 2024, 3.70);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (19, '123456789018', 'Mahasiswa 18', 'Informatika', 2024, 3.80);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (20, '123456789019', 'Mahasiswa 19', 'Sistem Informasi', 2024, 3.90);
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES (21, '123456789020', 'Mahasiswa 20', 'Informatika', 2024, 4.00);

-- 3. SEEDER COURSES
INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES ('MK001', 'Mata Kuliah 1', 2, 3, 35);
INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES ('MK002', 'Mata Kuliah 2', 3, 3, 40);
INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES ('MK003', 'Mata Kuliah 3', 2, 3, 45);
INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES ('MK004', 'Mata Kuliah 4', 3, 3, 50);
INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES ('MK005', 'Mata Kuliah 5', 2, 3, 55);
INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES ('MK006', 'Mata Kuliah 6', 3, 4, 60);
INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES ('MK007', 'Mata Kuliah 7', 2, 4, 65);
INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES ('MK008', 'Mata Kuliah 8', 3, 4, 70);
INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES ('MK009', 'Mata Kuliah 9', 2, 4, 75);
INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES ('MK0010', 'Mata Kuliah 10', 3, 4, 80);


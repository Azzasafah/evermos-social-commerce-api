
USE evermos_vix;

-- 1. Truncate / Clear old test data cleanly
SET FOREIGN_KEY_CHECKS = 0;
TRUNCATE TABLE detail_trx;
TRUNCATE TABLE log_produk;
TRUNCATE TABLE trx;
TRUNCATE TABLE foto_produk;
TRUNCATE TABLE produk;
TRUNCATE TABLE category;
TRUNCATE TABLE alamat;
TRUNCATE TABLE toko;
TRUNCATE TABLE users;
SET FOREIGN_KEY_CHECKS = 1;

-- 2. Insert Users (password: 'password123' hashed with bcrypt)
-- $2a$10$UoWp.Fp0p2aQ5X5M4f4Cpe4TsmJbJdD9g7E2k6fL8z3k3G0b1W2E2 is bcrypt for password123
INSERT INTO users (id, nama, kata_sandi, notelp, tanggal_lahir, jenis_kelamin, tentang, pekerjaan, email, id_provinsi, id_kota, is_admin, created_at, updated_at) VALUES
(1, 'Fajar Vibe', '$2a$10$eRUrhkaT2eTcCcS3q/jhHuruIAFBW.7c/zw3g1r0fFF3udaylr/Ty', '081234567890', '15/05/1998', 'Laki-laki', 'Founder & Master Distributor Evermos Ecosystem', 'Admin & Merchant', 'fajarvibe@example.com', '32', '3273', 1, NOW(), NOW()),
(2, 'Siti Reseller', '$2a$10$eRUrhkaT2eTcCcS3q/jhHuruIAFBW.7c/zw3g1r0fFF3udaylr/Ty', '089876543210', '10/08/1999', 'Perempuan', 'Top Reseller Bandung Raya busana & kecantikan halal', 'Reseller Pro', 'sitireseller@example.com', '32', '3273', 0, NOW(), NOW()),
(3, 'Ahmad Fauzi', '$2a$10$eRUrhkaT2eTcCcS3q/jhHuruIAFBW.7c/zw3g1r0fFF3udaylr/Ty', '081345678901', '22/03/1995', 'Laki-laki', 'Spesialis produk teknologi, audio, & aksesoris', 'Tech Reseller', 'ahmadfauzi@example.com', '31', '3171', 0, NOW(), NOW()),
(4, 'Nurul Hidayah', '$2a$10$eRUrhkaT2eTcCcS3q/jhHuruIAFBW.7c/zw3g1r0fFF3udaylr/Ty', '081567890123', '05/11/2001', 'Perempuan', 'Kurator modest fashion dan perlengkapan ibadah', 'Fashion Reseller', 'nurul@example.com', '35', '3578', 0, NOW(), NOW());

-- 3. Insert Tokos (One per user)
INSERT INTO toko (id, id_user, nama_toko, url_foto, created_at, updated_at) VALUES
(1, 1, 'Fajar Vibe Official Store', 'https://images.unsplash.com/photo-1441986300917-64674bd600d8?w=300&q=80', NOW(), NOW()),
(2, 2, 'Siti Hijab & Modest Wear', 'https://images.unsplash.com/photo-1472851294608-062f824d29cc?w=300&q=80', NOW(), NOW()),
(3, 3, 'Berkah Gadget & Living', 'https://images.unsplash.com/photo-1519389950473-47ba0277781c?w=300&q=80', NOW(), NOW()),
(4, 4, 'Nurul Syari Collection', 'https://images.unsplash.com/photo-1528698827591-e19ccd7bc23d?w=300&q=80', NOW(), NOW());

-- 4. Insert Categories
INSERT INTO category (id, nama_category, created_at, updated_at) VALUES
(1, 'Fashion Muslim & Hijab', NOW(), NOW()),
(2, 'Busana Pria Modern', NOW(), NOW()),
(3, 'Elektronik & Smart Gadget', NOW(), NOW()),
(4, 'Makanan & Minuman Halal', NOW(), NOW()),
(5, 'Aksesoris & Tas Premium', NOW(), NOW()),
(6, 'Perlengkapan Rumah Tangga', NOW(), NOW());

-- 5. Insert Alamat (Shipping Addresses)
INSERT INTO alamat (id, id_user, judul_alamat, nama_penerima, no_telp, detail_alamat, created_at, updated_at) VALUES
(1, 1, 'Kantor Pusat Evermos', 'Fajar Admin', '081234567890', 'Gedung Wisma Harmoni Lt. 4, Jl. Diponegoro No. 45, Bandung, Jawa Barat', NOW(), NOW()),
(2, 1, 'Rumah Pribadi Bandung', 'Fajar Pribadi', '081234567890', 'Komplek Dago Asri Blok C-12, Coblong, Kota Bandung', NOW(), NOW()),
(3, 2, 'Rumah Utama Siti', 'Siti Fatimah', '089876543210', 'Jl. Buah Batu No. 128, Lengkong, Kota Bandung', NOW(), NOW()),
(4, 2, 'Gudang Reseller Buahbatu', 'Siti Toko', '089876543210', 'Ruko Sentra Niaga Kav. 9, Batununggal, Kota Bandung', NOW(), NOW()),
(5, 3, 'Apartemen Jakarta', 'Ahmad Fauzi', '081345678901', 'Green Pramuka City Tower Chrysant 15A, Cempaka Putih, Jakarta Pusat', NOW(), NOW());

-- 6. Insert Products (Rich Catalog)
INSERT INTO produk (id, nama_produk, slug, harga_reseller, harga_konsumen, stok, deskripsi, id_toko, id_category, created_at, updated_at) VALUES
(1, 'Gamis Al-Zahra Silk Premium', 'gamis-al-zahra-silk-premium', 165000, 245000, 65, 'Gamis syari bahan sutra satin import dengan aksen renda bordir manual gold. Dingin dipakai, tidak nerawang, wudhu-friendly dengan zipper manset tangan.', 2, 1, NOW(), NOW()),
(2, 'Koko Kurta Modern Katun Toyobo', 'koko-kurta-modern-katun-toyobo', 110000, 175000, 80, 'Baju koko kurta pria bahan katun Toyobo fodu original. Jahitan garment rapi, kerah shanghai minimalis, dan saku paspol aktif.', 1, 2, NOW(), NOW()),
(3, 'Mukena Parasut Silk Traveling', 'mukena-parasut-silk-traveling', 85000, 135000, 120, 'Mukena mini pouch travel-friendly bahan parasut Korea silk ultra-ringan dan anti air. Mudah dilipat seukuran dompet tangan kecil.', 2, 1, NOW(), NOW()),
(4, 'Keyboard Mekanikal Wireless RGB', 'keyboard-mekanikal-wireless-rgb', 320000, 485000, 45, 'Keyboard mechanical 75% hot-swappable dengan 3 koneksi (Bluetooth 5.0, 2.4GHz Dongle, Type-C). Dilengkapi baterai 3000mAh tahan hingga 2 bulan.', 3, 3, NOW(), NOW()),
(5, 'Smartwatch AMOLED Fitness Tracker', 'smartwatch-amoled-fitness-tracker', 280000, 420000, 50, 'Layar AMOLED 1.43 inch Always-On Display dengan sensor detak jantung, SpO2, dan 100+ mode olahraga. Water resistant 5 ATM.', 3, 3, NOW(), NOW()),
(6, 'Madu Hutan Murni Baduy 500gr', 'madu-hutan-murni-baduy-500gr', 75000, 120000, 95, 'Madu mentah alami tanpa proses pasteurisasi dari hutan Ujung Kulon Baduy. Kaya antioksidan, enzim alami, dan bersertifikat Halal MUI.', 1, 4, NOW(), NOW()),
(7, 'Tas Ransel Canvas Waterproof Urban', 'tas-ransel-canvas-waterproof-urban', 145000, 230000, 40, 'Backpack multifungsi dengan kompartemen laptop 15.6 inch busa tebal, material kanvas cordura anti-air, dan port USB charger eksternal.', 1, 5, NOW(), NOW()),
(8, 'Set Panci Granit Keramik Anti Lengket', 'set-panci-granit-keramik-anti-lengket', 240000, 375000, 30, 'Set 4 in 1 wajan frypan & casserole induksi coating batu granit Jerman. Bebas PFOA/PTFE, mudah dibersihkan dan hemat minyak goreng.', 1, 6, NOW(), NOW()),
(9, 'Hijab Pashmina Ceruty Babydoll', 'hijab-pashmina-ceruty-babydoll', 28000, 50000, 250, 'Pashmina ceruty premium flowy, jatuh sempurna saat di-styling, dan tidak mudah kusut. Pilihan warna earth tone terlengkap.', 4, 1, NOW(), NOW()),
(10, 'Kemeja Flannel Tartan Slimfit', 'kemeja-flannel-tartan-slimfit', 105000, 165000, 70, 'Kemeja pria motif kotak klasik bahan katun flannel premium tebal berbulu halus. Nyaman dipakai casual harian maupun semi-formal.', 1, 2, NOW(), NOW()),
(11, 'Earphone TWS Bluetooth 5.3 ANC', 'earphone-tws-bluetooth-53-anc', 150000, 240000, 85, 'Earphone wireless low-latency dengan Active Noise Cancelling dan 4 mikrofon jernih untuk panggilan telepon suara HD.', 3, 3, NOW(), NOW()),
(12, 'Kurma Sukari Al-Qassim King 1kg', 'kurma-sukari-al-qassim-king-1kg', 65000, 105000, 110, 'Kurma basah daging lembut manis alami langsung dari perkebunan Al-Qassim Arab Saudi. Segar disimpan di chiller bersuhu stabil.', 1, 4, NOW(), NOW());

-- 7. Insert Foto Produk (High-res curated images)
INSERT INTO foto_produk (id_produk, url, created_at, updated_at) VALUES
(1, 'https://images.unsplash.com/photo-1585487000160-6ebcfceb0d03?w=800&q=80', NOW(), NOW()),
(2, 'https://images.unsplash.com/photo-1602810318383-e386cc2a3ccf?w=800&q=80', NOW(), NOW()),
(3, 'https://images.unsplash.com/photo-1596755094514-f87e34085b2c?w=800&q=80', NOW(), NOW()),
(4, 'https://images.unsplash.com/photo-1587829741301-dc798b83add3?w=800&q=80', NOW(), NOW()),
(5, 'https://images.unsplash.com/photo-1523275335684-37898b6baf30?w=800&q=80', NOW(), NOW()),
(6, 'https://images.unsplash.com/photo-1587049352846-4a222e784d38?w=800&q=80', NOW(), NOW()),
(7, 'https://images.unsplash.com/photo-1553062407-98eeb64c6a62?w=800&q=80', NOW(), NOW()),
(8, 'https://images.unsplash.com/photo-1584269600464-37b1b58a9fe7?w=800&q=80', NOW(), NOW()),
(9, 'https://images.unsplash.com/photo-1609357605129-26f69add5d6e?w=800&q=80', NOW(), NOW()),
(10, 'https://images.unsplash.com/photo-1603252109303-2751441dd157?w=800&q=80', NOW(), NOW()),
(11, 'https://images.unsplash.com/photo-1590658268037-6bf12165a8df?w=800&q=80', NOW(), NOW()),
(12, 'https://images.unsplash.com/photo-1579783900882-c0d3dad7b119?w=800&q=80', NOW(), NOW());

-- 8. Insert Sample Transactions (Trx, LogProduk, DetailTrx)
-- Trx 1
INSERT INTO trx (id, id_user, alamat_pengiriman, harga_total, kode_invoice, method_bayar, created_at, updated_at) VALUES
(1, 1, 1, 490000, 'INV-1711289001', 'bca', DATE_SUB(NOW(), INTERVAL 2 DAY), DATE_SUB(NOW(), INTERVAL 2 DAY)),
(2, 2, 3, 420000, 'INV-1711345002', 'qris', DATE_SUB(NOW(), INTERVAL 1 DAY), DATE_SUB(NOW(), INTERVAL 1 DAY));

-- Log Produk Snapshot for Trx 1
INSERT INTO log_produk (id, id_produk, nama_produk, slug, harga_reseller, harga_konsumen, deskripsi, id_toko, id_category, created_at, updated_at) VALUES
(1, 1, 'Gamis Al-Zahra Silk Premium', 'gamis-al-zahra-silk-premium', 165000, 245000, 'Snapshot gamis Al-Zahra saat transaksi', 2, 1, DATE_SUB(NOW(), INTERVAL 2 DAY), DATE_SUB(NOW(), INTERVAL 2 DAY)),
(2, 5, 'Smartwatch AMOLED Fitness Tracker', 'smartwatch-amoled-fitness-tracker', 280000, 420000, 'Snapshot smartwatch saat transaksi', 3, 3, DATE_SUB(NOW(), INTERVAL 1 DAY), DATE_SUB(NOW(), INTERVAL 1 DAY));

-- Detail Trx
INSERT INTO detail_trx (id, id_trx, id_log_produk, id_toko, kuantitas, harga_total, created_at, updated_at) VALUES
(1, 1, 1, 2, 2, 490000, DATE_SUB(NOW(), INTERVAL 2 DAY), DATE_SUB(NOW(), INTERVAL 2 DAY)),
(2, 2, 2, 3, 1, 420000, DATE_SUB(NOW(), INTERVAL 1 DAY), DATE_SUB(NOW(), INTERVAL 1 DAY));

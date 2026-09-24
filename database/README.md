# Database & Seed Management

This directory contains initialization and database seeding scripts for the **`evermos_vix`** database.

## Directory Contents

- **`seeds/seed.sql`**: Pure SQL script for MySQL/MariaDB. It cleans existing tables (`TRUNCATE`) with foreign key checks toggled, and seeds demo data including Users (Admin & Resellers), Stores (Toko), Product Categories, Shipping Addresses, rich Product Catalogs with Unsplash photos, Completed Transactions, and `log_produk` audit trails.
- **`seeds/seed_data.py`**: Automated Python runner to execute `seed.sql` directly into MySQL CLI (`mysql -u root evermos_vix`).

---

## How to Run Seed Data

### Method 1: Using MySQL CLI (Recommended)
```bash
mysql -u root -p evermos_vix < database/seeds/seed.sql
```
*(If your local MySQL has no password set, omit `-p`)*:
```bash
mysql -u root evermos_vix < database/seeds/seed.sql
```

### Method 2: Using Python Automation Runner
Ensure Python 3 is installed and `mysql` is available in your system's PATH:
```bash
python database/seeds/seed_data.py
```

### Method 3: Using GUI (DBeaver / phpMyAdmin / MySQL Workbench)
1. Open the `evermos_vix` database.
2. Open the file `database/seeds/seed.sql` in your SQL Editor.
3. Execute the script.

---

## Seeded Demo Accounts

All demo accounts share the default password: **`password123`**

| Role | Phone Number (Login) | Password | Store Name / Description |
| :--- | :--- | :--- | :--- |
| **Admin** | `081234567890` | `password123` | Fajar Vibe Official Store (Master Distributor) |
| **Reseller Pro** | `089876543210` | `password123` | Siti Hijab & Modest Wear (Bandung) |
| **Reseller Tech** | `081345678901` | `password123` | Berkah Gadget & Living (Jakarta) |
| **Reseller Fashion**| `081567890123` | `password123` | Nurul Syari Collection (Surabaya) |

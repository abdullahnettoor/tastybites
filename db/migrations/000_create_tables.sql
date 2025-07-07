

-- TastyBites Database Schema and Sample Data
-- This file creates all tables, indexes, triggers, and sample data

-- =============================================================================
-- UTILITY FUNCTIONS
-- =============================================================================

-- Create trigger function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- =============================================================================
-- USERS TABLE
-- =============================================================================

-- Create users table
CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'user' CHECK (role IN ('user', 'admin', 'manager')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes on users table
CREATE INDEX IF NOT EXISTS idx_users_email ON public.users(email);

-- Create trigger for users table
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON public.users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- =============================================================================
-- MENU ITEMS TABLE
-- =============================================================================

-- Create menu_items table
CREATE TABLE IF NOT EXISTS public.menu_items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price INTEGER NOT NULL CHECK (price > 0), -- Price in cents to avoid decimal issues
    category VARCHAR(50) NOT NULL,
    image_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes on menu_items table
CREATE INDEX IF NOT EXISTS idx_menu_items_category ON public.menu_items(category);
CREATE INDEX IF NOT EXISTS idx_menu_items_name ON public.menu_items(name);

-- Create trigger for menu_items table
CREATE TRIGGER update_menu_items_updated_at
    BEFORE UPDATE ON public.menu_items
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- =============================================================================
-- TABLES TABLE
-- =============================================================================

-- Create tables table
CREATE TABLE IF NOT EXISTS public.tables (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    seats INTEGER NOT NULL CHECK (seats > 0),
    status VARCHAR(20) DEFAULT 'available' CHECK (status IN ('available', 'reserved')),
    booked_by INTEGER REFERENCES public.users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes on tables table
CREATE INDEX IF NOT EXISTS idx_tables_status ON public.tables(status);
CREATE INDEX IF NOT EXISTS idx_tables_booked_by ON public.tables(booked_by);

-- Create trigger for tables table
CREATE TRIGGER update_tables_updated_at
    BEFORE UPDATE ON public.tables
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- =============================================================================
-- ORDERS TABLE
-- =============================================================================

-- Create orders table
CREATE TABLE IF NOT EXISTS public.orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    table_id INTEGER REFERENCES public.tables(id) ON DELETE SET NULL,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'cancelled')),
    total_price INTEGER NOT NULL DEFAULT 0 CHECK (total_price >= 0), -- Price in cents
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes on orders table
CREATE INDEX IF NOT EXISTS idx_orders_user_id ON public.orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_table_id ON public.orders(table_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON public.orders(status);
CREATE INDEX IF NOT EXISTS idx_orders_created_at ON public.orders(created_at);

-- Create trigger for orders table
CREATE TRIGGER update_orders_updated_at
    BEFORE UPDATE ON public.orders
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- =============================================================================
-- ORDER ITEMS TABLE
-- =============================================================================

-- Create order_items table (junction table for orders and menu_items)
CREATE TABLE IF NOT EXISTS public.order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES public.orders(id) ON DELETE CASCADE,
    menu_item_id INTEGER NOT NULL REFERENCES public.menu_items(id) ON DELETE RESTRICT,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    price INTEGER NOT NULL CHECK (price > 0), -- Price per item in cents at time of order
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes on order_items table
CREATE INDEX IF NOT EXISTS idx_order_items_order_menu ON public.order_items(order_id, menu_item_id);
CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON public.order_items(order_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_order_items_unique ON public.order_items(order_id, menu_item_id);

-- =============================================================================
-- SAMPLE DATA
-- =============================================================================

-- Insert sample users (using bcrypt hashed passwords)
-- Default password for all users: "password123"
INSERT INTO public.users (name, email, password, role) VALUES
('Admin', 'admin@tastybites.com', '$2a$10$85UC/kMB3jsoINLoDV5dk.8uMlYrQHhxbV.21SLnwRxjgKBcZcu1u', 'admin'),
('John Doe', 'john@example.com', '$2a$10$85UC/kMB3jsoINLoDV5dk.8uMlYrQHhxbV.21SLnwRxjgKBcZcu1u', 'user'),
('Jane Smith', 'jane@example.com', '$2a$10$85UC/kMB3jsoINLoDV5dk.8uMlYrQHhxbV.21SLnwRxjgKBcZcu1u', 'user'),
('Manager', 'manager@tastybites.com', '$2a$10$85UC/kMB3jsoINLoDV5dk.8uMlYrQHhxbV.21SLnwRxjgKBcZcu1u', 'admin')
ON CONFLICT (email) DO NOTHING;

-- Insert sample menu items
INSERT INTO public.menu_items (name, description, price, category, image_url) VALUES
('Margherita Pizza', 'Classic pizza with tomato sauce, mozzarella, and fresh basil', 1299, 'Pizza', 'https://images.unsplash.com/photo-1574071318508-1cdbab80d002'),
('Chicken Caesar Salad', 'Crisp romaine lettuce with grilled chicken, parmesan, and caesar dressing', 1149, 'Salads', 'https://images.unsplash.com/photo-1546793665-c74683f339c1'),
('Beef Burger', 'Juicy beef patty with lettuce, tomato, cheese, and special sauce', 1499, 'Burgers', 'https://images.unsplash.com/photo-1568901346375-23c9450c58cd'),
('Chocolate Brownie', 'Rich chocolate brownie served with vanilla ice cream', 699, 'Desserts', 'https://images.unsplash.com/photo-1606313564200-e75d5e30476c'),
('Pepperoni Pizza', 'Classic pepperoni pizza with mozzarella cheese', 1499, 'Pizza', 'https://images.unsplash.com/photo-1565299624946-b28f40a0ca4b'),
('Greek Salad', 'Fresh vegetables with feta cheese and olive oil dressing', 999, 'Salads', 'https://images.unsplash.com/photo-1540420773420-3366772f4999'),
('Fish & Chips', 'Beer-battered cod with crispy fries and tartar sauce', 1699, 'Main Course', 'https://images.unsplash.com/photo-1544025162-d76694265947'),
('Tiramisu', 'Classic Italian dessert with coffee-soaked ladyfingers', 799, 'Desserts', 'https://images.unsplash.com/photo-1571877227200-a0d98ea607e9'),
('BBQ Chicken Wings', 'Smoky BBQ chicken wings with celery sticks', 1299, 'Appetizers', 'https://images.unsplash.com/photo-1608039755401-742074f0548d'),
('Vegetarian Pasta', 'Penne pasta with seasonal vegetables in garlic olive oil', 1199, 'Pasta', 'https://images.unsplash.com/photo-1621996346565-e3dbc1d56d0e')
ON CONFLICT DO NOTHING;

-- Insert sample tables
INSERT INTO public.tables (name, seats, status, booked_by) VALUES
('Table 1', 4, 'available', NULL),
('Table 2', 2, 'available', NULL),
('Table 3', 6, 'reserved', 2), -- Reserved by john_doe
('Table 4', 4, 'available', NULL),
('Table 5', 8, 'reserved', 4), -- Reserved by manager
('Table 6', 2, 'available', NULL),
('Table 7', 4, 'available', NULL),
('Table 8', 6, 'reserved', 3), -- Reserved by jane_smith
('Bar Counter', 6, 'available', NULL),
('Private Room', 12, 'available', NULL)
ON CONFLICT (name) DO NOTHING;

-- Insert sample orders
INSERT INTO public.orders (user_id, table_id, status, total_price) VALUES
(2, 1, 'completed', 2448), -- John's order at Table 1: Margherita Pizza + Chicken Caesar Salad
(3, 2, 'pending', 1499),   -- Jane's order at Table 2: Beef Burger
(2, 4, 'pending', 2198),   -- John's another order at Table 4: Pepperoni Pizza + Chocolate Brownie
(4, 9, 'completed', 2698)  -- Manager's order at Bar Counter: Fish & Chips + BBQ Wings
ON CONFLICT DO NOTHING;

-- Insert sample order items
INSERT INTO public.order_items (order_id, menu_item_id, quantity, price) VALUES
-- John's first order (order_id: 1)
(1, 1, 1, 1299), -- 1x Margherita Pizza
(1, 2, 1, 1149), -- 1x Chicken Caesar Salad

-- Jane's order (order_id: 2)
(2, 3, 1, 1499), -- 1x Beef Burger

-- John's second order (order_id: 3)
(3, 5, 1, 1499), -- 1x Pepperoni Pizza
(3, 4, 1, 699),  -- 1x Chocolate Brownie

-- Manager's order (order_id: 4)
(4, 7, 1, 1699), -- 1x Fish & Chips
(4, 9, 1, 999)   -- 1x BBQ Chicken Wings
ON CONFLICT DO NOTHING;

-- =============================================================================
-- COMPLETION MESSAGE
-- =============================================================================

-- Display completion message
DO $$
BEGIN
    RAISE NOTICE 'TastyBites database schema and sample data created successfully!';
    RAISE NOTICE 'Sample users created with password: "password123"';
    RAISE NOTICE 'Admin user: admin@tastybites.com';
    RAISE NOTICE 'Regular users: john@example.com, jane@example.com';
    RAISE NOTICE 'Manager user: manager@tastybites.com';
END $$;
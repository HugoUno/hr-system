-- 創建員工表
CREATE TABLE IF NOT EXISTS employees (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    employee_id VARCHAR(20) UNIQUE,
    name VARCHAR(50),
    position VARCHAR(50),
    department VARCHAR(50),
    level INT DEFAULT 1,
    salary INT DEFAULT 0,
    email VARCHAR(100),
    phone VARCHAR(20),
    address VARCHAR(200),
    join_date DATETIME,
    leave_balance INT DEFAULT 0,
    status VARCHAR(20),
    INDEX idx_employee_id (employee_id),
    INDEX idx_deleted_at (deleted_at)
);

-- 創建請假記錄表
CREATE TABLE IF NOT EXISTS leave_requests (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    employee_id BIGINT UNSIGNED,
    leave_type VARCHAR(20),
    start_date DATETIME,
    end_date DATETIME,
    duration FLOAT,
    reason VARCHAR(500),
    status VARCHAR(20),
    approved_by BIGINT UNSIGNED,
    approved_at DATETIME,
    approver_note VARCHAR(500),
    INDEX idx_employee_id (employee_id),
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    FOREIGN KEY (approved_by) REFERENCES employees(id)
); 
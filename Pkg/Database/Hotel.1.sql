-- Usuario
CREATE TABLE Usuario (
    id_usuario INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100),
    apellido VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    telefono VARCHAR(20),
    direccion VARCHAR(255)
);

-- Habitación
CREATE TABLE Habitacion (
    id_habitacion INT AUTO_INCREMENT PRIMARY KEY,
    numero VARCHAR(10),
    tipo ENUM('sencilla','doble','suite'),
    precio DECIMAL(10, 2),
    estado ENUM('disponible', 'reservada', 'ocupada', 'no disponible')
);

-- Reserva
CREATE TABLE Reserva (
    id_reserva INT AUTO_INCREMENT PRIMARY KEY,
    id_usuario INT,
    id_habitacion INT,
    fecha_reserva DATE,
    estado ENUM('confirmada', 'cancelada'),
    FOREIGN KEY (id_usuario) REFERENCES Usuario(id_usuario),
    FOREIGN KEY (id_habitacion) REFERENCES Habitacion(id_habitacion)
);

-- Pago
CREATE TABLE Pago (
    id_pago INT AUTO_INCREMENT PRIMARY KEY,
    id_reserva INT,
    monto DECIMAL(10, 2),
    metodo_pago ENUM('tarjeta','efectivo'),
    fecha_pago DATE,
    FOREIGN KEY (id_reserva) REFERENCES Reserva(id_reserva)
);

-- Llave
CREATE TABLE Llave (
    id_llave INT AUTO_INCREMENT PRIMARY KEY,
    id_reserva INT,
    id_habitacion INT,
    tipo_llave ENUM('fisica', 'electronica'),
    estado_llave ENUM('activa', 'desactivada'),
    FOREIGN KEY (id_reserva) REFERENCES Reserva(id_reserva),
    FOREIGN KEY (id_habitacion) REFERENCES Habitacion(id_habitacion)
);

-- Check-in/Check-out
CREATE TABLE CheckInCheckOut (
    id_checkin INT AUTO_INCREMENT PRIMARY KEY,
    id_reserva INT,
    fecha_checkin DATE,
    fecha_checkout DATE,
    FOREIGN KEY (id_reserva) REFERENCES Reserva(id_reserva)
);

-- Personal
CREATE TABLE Personal (
    id_personal INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100),
    apellido VARCHAR(100),
    rol ENUM('recepcionista','limpiesa','gerente', 'botones', 'supervisor')
);

-- Factura Electrónica
CREATE TABLE FacturaElectronica (
    id_factura INT AUTO_INCREMENT PRIMARY KEY,
    id_pago INT,
    fecha_factura DATE,
    total DECIMAL(10, 2),
    FOREIGN KEY (id_pago) REFERENCES Pago(id_pago)
);

-- Tabla intermedia para la relación muchos a muchos entre Personal y Habitacion
CREATE TABLE Personal_Habitacion (
    id_personal_habitacion INT AUTO_INCREMENT PRIMARY KEY,
    id_personal INT,
    id_habitacion INT,
    fecha_asignacion DATE,
    FOREIGN KEY (id_personal) REFERENCES Personal(id_personal),
    FOREIGN KEY (id_habitacion) REFERENCES Habitacion(id_habitacion)
);
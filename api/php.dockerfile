# Use the official PHP image with Apache
FROM php:8.1-apache

# Install PHP extensions
RUN docker-php-ext-install mysqli pdo pdo_mysql

# Enable Apache mod_rewrite
RUN a2enmod rewrite

# Set working directory
WORKDIR /var/www/html

# Copy your application code (optional)
# COPY . /var/www/html/

# Expose port 80
EXPOSE 80

class Product < ApplicationRecord
  validates :name,  length: {minimum:4}
  validates :price, presence: true, format: { with: /\A\d+(?:\.\d{2})?\z/ }, numericality: { greater_than: 0, less_than: 1000000}
  validates :brand, :amount, :description, presence: true
  validates :amount, numericality: { only_numeric: true }
end

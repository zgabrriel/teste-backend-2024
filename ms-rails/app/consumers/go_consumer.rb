# frozen_string_literal: true
class GoConsumer < ApplicationConsumer
  def consume
    messages.each do |message|
      data = message.payload

      existing_product = Product.find_by(id: data["id"])

      if existing_product.nil?
        product = Product.new(
          id: data["id"],
          name: data["name"],
          brand: data["brand"],
          price: data["price"],
          description: data["description"],
          created_at: data["created_at"],
          updated_at: data["updated_at"],
          amount: data["amount"]
        )

        product.save!
      else
        existing_product.update(data)
      end
    end
  end
end

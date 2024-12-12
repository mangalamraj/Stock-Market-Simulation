"use client";

import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { useState } from "react";

interface BuyStockFormProps {
  onPurchase?: () => void;
  selectedStock: string;
}

const BuyStockForm = ({ onPurchase, selectedStock }: BuyStockFormProps) => {
  const [quantity, setQuantity] = useState<number>(0);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const email = localStorage.getItem("email");
    if (!email) {
      return;
    }

    if (quantity <= 0) {
      return;
    }

    setLoading(true);
    try {
      const response = await fetch("http://localhost:8080/stock/buy", {
        method: "POST",
        body: JSON.stringify({ 
          email, 
          stock: selectedStock, 
          quantity 
        }),
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (!response.ok) {
        const error = await response.text();
        throw new Error(error);
      }

      const data = await response.json();
      console.log(data);
      setQuantity(0);
      
      if (onPurchase) {
        onPurchase();
      }
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="bg-black text-white pb-10">
      <div className="flex flex-col gap-4 container mx-auto text-black border border-gray-300 rounded-md p-4 w-1/2">
        <div className="text-2xl font-bold text-white">
          Buy {selectedStock} Stock
        </div>
        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
          <Input
            type="number"
            placeholder="Enter the quantity"
            value={quantity}
            onChange={(e) => setQuantity(Number(e.target.value))}
            min="1"
          />
          <Button type="submit" disabled={loading}>
            {loading ? "Buying..." : "Buy"}
          </Button>
        </form>
      </div>
    </div>
  );
};

export default BuyStockForm;

"use client";

import BuyStockForm from "@/components/buyStockForm/buyStockForm";
import StockGraph from "@/components/stockGraph/stockgraph";
import { useState } from "react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const STOCKS = [
  { value: "AAPL", label: "Apple Inc." },
  { value: "GOOG", label: "Google" },
  { value: "AMZN", label: "Amazon" },
  { value: "TSLA", label: "Tesla" },
];

export default function BuyStocks() {
  const [refreshTrigger, setRefreshTrigger] = useState(0);
  const [selectedStock, setSelectedStock] = useState("AAPL");

  const handlePurchase = () => {
    setRefreshTrigger((prev) => prev + 1);
  };

  return (
    <div className="bg-black min-h-screen p-4">
      <div className="container mx-auto">
        <div className="mb-10 flex gap-4 items-center justify-center align-middle">
          <h2 className="text-white text-2xl">Select Stock</h2>
          <Select value={selectedStock} onValueChange={setSelectedStock}>
            <SelectTrigger className="w-[200px]">
              <SelectValue placeholder="Select a stock" />
            </SelectTrigger>
            <SelectContent>
              {STOCKS.map((stock) => (
                <SelectItem key={stock.value} value={stock.value}>
                  {stock.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>

        <div className="grid gap-8">
          <StockGraph
            stock={selectedStock}
            key={`${selectedStock}-${refreshTrigger}`}
            shouldRefresh={true}
          />
          <BuyStockForm
            onPurchase={handlePurchase}
            selectedStock={selectedStock}
          />
        </div>
      </div>
    </div>
  );
}

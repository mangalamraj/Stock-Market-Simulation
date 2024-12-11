"use client";

import StockGraph from "@/components/stockGraph/stockgraph";
import { useState } from "react";

const Buystocks = () => {
  const [stock, setStock] = useState("AAPL");

  const handleStockChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setStock(e.target.value);
  };

  return (
    <div className="bg-black text-white pt-10">
      <div className="flex items-center justify-center gap-4 pb-4">
        <div className="flex gap-2 items-center justify-center">
          <input
            type="radio"
            id="AAPL"
            name="stock"
            value="AAPL"
            checked={stock === "AAPL"}
            onChange={handleStockChange}
          />
          <label htmlFor="AAPL">AAPL</label>
        </div>
        <div className="flex gap-2 items-center justify-center">
          <input
            type="radio"
            id="GOOGL"
            name="stock"
            value="GOOG"
            checked={stock === "GOOG"}
            onChange={handleStockChange}
          />

          <label htmlFor="GOOG">GOOG</label>
        </div>
        <div className="flex gap-2 items-center justify-center">
          <input
            type="radio"
            id="AMZN"
            name="stock"
            value="AMZN"
            checked={stock === "AMZN"}
            onChange={handleStockChange}
          />
          <label htmlFor="AMZN">AMZN</label>
        </div>
        <div className="flex gap-2 items-center justify-center">
          <input
            type="radio"
            id="AMZN"
            name="stock"
            value="TSLA"
            checked={stock === "TSLA"}
            onChange={handleStockChange}
          />
          <label htmlFor="TSLA">TSLA</label>
        </div>
      </div>

      <StockGraph stock={stock} />
    </div>
  );
};

export default Buystocks;

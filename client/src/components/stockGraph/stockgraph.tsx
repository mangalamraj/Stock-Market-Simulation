"use client";
import React, { useEffect, useState, useCallback } from "react";
import { Line } from "react-chartjs-2";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";

// Register ChartJS components
ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

interface StockEntry {
  timestamp: number;
  price: number;
}

interface ChartData {
  labels: string[];
  datasets: {
    label: string;
    data: number[];
    fill: boolean;
    borderColor: string;
    tension: number;
  }[];
}

function StockGraph({
  stock,
  shouldRefresh = false,
}: {
  stock: string;
  shouldRefresh?: boolean;
}) {
  const [chartData, setChartData] = useState<ChartData | null>(null);

  const fetchData = useCallback(async () => {
    try {
      const response = await fetch(
        `http://localhost:8080/stock/history?stock=${stock}`
      );
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data: StockEntry[] = await response.json();

      const timestamps = data.map((entry: StockEntry) =>
        new Date(entry.timestamp * 1000).toLocaleString()
      );
      const prices = data.map((entry: StockEntry) => entry.price);

      setChartData({
        labels: timestamps,
        datasets: [
          {
            label: `${stock} Price`,
            data: prices,
            fill: false,
            borderColor: "rgba(75,192,192,1)",
            tension: 0.1,
          },
        ],
      });
    } catch (error) {
      console.error("Error fetching stock history:", error);
    }
  }, [stock]);

  useEffect(() => {
    fetchData();

    // If shouldRefresh is true, set up an interval to refresh data
    if (shouldRefresh) {
      const interval = setInterval(fetchData, 5000); // Refresh every 5 seconds
      return () => clearInterval(interval);
    }
  }, [stock, shouldRefresh, fetchData]);

  if (!chartData) return <div>Loading...</div>;

  return (
    <div className="bg-black text-white pb-10">
      <div className="container m-auto text-center">
        <h2>{stock} Stock Price History</h2>
        <div className="w-full h-[600px] flex justify-center items-center">
          <Line data={chartData} />
        </div>
      </div>
    </div>
  );
}

export default StockGraph;

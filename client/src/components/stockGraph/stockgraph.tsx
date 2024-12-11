"use client";
import React, { useEffect, useState } from "react";
import { Line } from "react-chartjs-2";
import {
  Chart,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";

// Register the necessary components
Chart.register(
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

function StockGraph({ stock }: { stock: string }) {
  const [chartData, setChartData] = useState<ChartData | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(
          `http://localhost:8080/stock/history?stock=${stock}`,
          {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
            },
          }
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
    };

    fetchData();
  }, [stock]);

  if (!chartData) return <div>Loading...</div>;

  return (
    <div className="bg-black text-white h-screen">
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

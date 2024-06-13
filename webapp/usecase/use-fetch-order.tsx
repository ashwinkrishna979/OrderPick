import { useState, useEffect } from 'react';

// Define the shape of the order data
interface Order {
  order_id: string;
  item_id: string;
  created_at: string;
  packing_status: boolean;
}

// Define the return type for the hook
interface UseFetchOrdersReturn {
  data: Order[];
  loading: boolean;
  error: Error | null;
}

const useFetchOrders = (token: string): UseFetchOrdersReturn => {
  const [data, setData] = useState<Order[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8000/orders', {
            
        method: 'GET',
          headers: {
            'token': `${token}`,
            'Content-Type': 'application/json'
          }
        });
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const result: Order[] = await response.json();
        setData(result);
      } catch (error) {
        setError(error as Error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [token]);

  return { data, loading, error };
};

export default useFetchOrders;

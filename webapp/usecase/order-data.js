import React from "react";
import { v4 as uuidv4 } from 'uuid';

// Columns definition for Orders table
const columns = [
  { name: "Order ID", uid: "order_id", sortable: true },
  { name: "Item ID", uid: "item_id", sortable: true },
  { name: "Created At", uid: "created_at", sortable: true },
  { name: "Packing Status", uid: "packing_status", sortable: true },
];

// Sample data for Orders
const orders = [
  {
    order_id: uuidv4(),
    item_id: "12345",
    created_at: new Date("2023-01-01T10:00:00Z"),
    packing_status: true,
  },
  {
    order_id: uuidv4(),
    item_id: "67890",
    created_at: new Date("2023-02-15T14:30:00Z"),
    packing_status: false,
  },
  {
    order_id: uuidv4(),
    item_id: "11223",
    created_at: new Date("2023-03-20T09:15:00Z"),
    packing_status: true,
  },
  {
    order_id: uuidv4(),
    item_id: "44556",
    created_at: new Date("2023-04-05T17:45:00Z"),
    packing_status: false,
  },
  {
    order_id: uuidv4(),
    item_id: "78901",
    created_at: new Date("2023-05-25T08:00:00Z"),
    packing_status: true,
  },
  {
    order_id: uuidv4(),
    item_id: "23456",
    created_at: new Date("2023-06-10T11:30:00Z"),
    packing_status: false,
  },
  {
    order_id: uuidv4(),
    item_id: "34567",
    created_at: new Date("2023-07-15T16:00:00Z"),
    packing_status: true,
  },
  {
    order_id: uuidv4(),
    item_id: "45678",
    created_at: new Date("2023-08-01T13:45:00Z"),
    packing_status: false,
  },
  {
    order_id: uuidv4(),
    item_id: "56789",
    created_at: new Date("2023-09-12T07:30:00Z"),
    packing_status: true,
  },
  {
    order_id: uuidv4(),
    item_id: "67890",
    created_at: new Date("2023-10-18T19:15:00Z"),
    packing_status: false,
  },
];

export { columns, orders };

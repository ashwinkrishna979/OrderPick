import React from "react";

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
    order_id: "e1b1b6c8-6a23-4b6a-b3ec-abcde1234567",
    item_id: "12345",
    created_at: "2023-01-01T10:00:00Z",
    packing_status: true,
  },
  {
    order_id: "f2b2c6d9-7b34-4c7a-c4ec-bcdef2345678",
    item_id: "67890",
    created_at: "2023-02-15T14:30:00Z",
    packing_status: false,
  },
  {
    order_id: "g3c3d7e1-8c45-4d8a-d5ec-cdef3456789a",
    item_id: "11223",
    created_at: "2023-03-20T09:15:00Z",
    packing_status: true,
  },
  {
    order_id: "h4d4e8f2-9d56-4e9a-e6ec-def456789ab1",
    item_id: "44556",
    created_at: "2023-04-05T17:45:00Z",
    packing_status: false,
  },
  {
    order_id: "i5e5f9g3-0e67-4f0a-f7ec-ef567890bc12",
    item_id: "78901",
    created_at: "2023-05-25T08:00:00Z",
    packing_status: true,
  },
  {
    order_id: "j6f6g0h4-1f78-401a-g8ec-f678901cd123",
    item_id: "23456",
    created_at: "2023-06-10T11:30:00Z",
    packing_status: false,
  },
  {
    order_id: "k7g7h1i5-2g89-402a-h9ec-g789012de234",
    item_id: "34567",
    created_at: "2023-07-15T16:00:00Z",
    packing_status: true,
  },
  {
    order_id: "l8h8i2j6-3h90-403a-i0ec-h890123ef345",
    item_id: "45678",
    created_at: "2023-08-01T13:45:00Z",
    packing_status: false,
  },
  {
    order_id: "m9i9j3k7-4i01-404a-j1ec-i901234fg456",
    item_id: "56789",
    created_at: "2023-09-12T07:30:00Z",
    packing_status: true,
  },
  {
    order_id: "n0j0k4l8-5j12-405a-k2ec-j012345gh567",
    item_id: "67890",
    created_at: "2023-10-18T19:15:00Z",
    packing_status: false,
  },
];

export { columns, orders };

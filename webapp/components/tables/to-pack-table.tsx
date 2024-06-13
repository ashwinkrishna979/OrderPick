'use client'
import React, { useState, useMemo, useCallback, useEffect } from "react";
import {
  Table,
  TableHeader,
  TableColumn,
  TableBody,
  TableRow,
  TableCell,
  Input,
  Button,
  Pagination,
  Selection,
  SortDescriptor
} from "@nextui-org/react";
import { PlusIcon } from "../icons";
import { SearchIcon } from "../icons";
import { columns } from "../../usecase/order-data";
import useFetchOrders from '../../usecase/use-fetch-order';

const INITIAL_VISIBLE_COLUMNS = ["order_id", "item_id", "created_at", "packing_status"];

const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X25hbWUiOiIiLCJMYXN0X25hbWUiOiIiLCJVaWQiOiIiLCJleHAiOjE3MTg4NDE5NTF9.6cNzjUfp6KDAVEO2Oc-VedZhSV6FpuUx4DWLp5tG9ao';

export default function PackedTable() {
  var { data: orders, loading, error } = useFetchOrders(token);

  const [filterValue, setFilterValue] = useState("");
  const [selectedKeys, setSelectedKeys] = useState<Selection>(new Set([]));
  const [visibleColumns] = useState<Selection>(new Set(INITIAL_VISIBLE_COLUMNS));
  const [rowsPerPage, setRowsPerPage] = useState(5);
  const [sortDescriptor, setSortDescriptor] = useState<SortDescriptor>({
    column: "order_id",
    direction: "ascending",
  });
  const [page, setPage] = useState(1);

  useEffect(() => {
    console.log("orders:", orders);
  }, [orders]);

  const pages = Math.ceil(orders.length / rowsPerPage);

  const hasSearchFilter = Boolean(filterValue);

  const headerColumns = useMemo(() => {
    if (visibleColumns === "all") return columns;

    return columns.filter((column) => Array.from(visibleColumns).includes(column.uid));
  }, [visibleColumns]);

  const filteredItems = useMemo(() => {
    if (!Array.isArray(orders)){
      orders=[];
    }
    let filteredOrders = [...orders];

    if (hasSearchFilter) {
      filteredOrders = filteredOrders.filter((order) =>
        order.order_id.toLowerCase().includes(filterValue.toLowerCase()),
      );
    }

    return filteredOrders;
  }, [filterValue, orders, hasSearchFilter]);

  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;

    return filteredItems.slice(start, end);
  }, [page, filteredItems, rowsPerPage]);

  const sortedItems = useMemo(() => {
    return [...items].sort((a: typeof orders[0], b: typeof orders[0]) => {
      const first = a[sortDescriptor.column as keyof typeof orders[0]] as string;
      const second = b[sortDescriptor.column as keyof typeof orders[0]] as string;
      const cmp = first < second ? -1 : first > second ? 1 : 0;

      return sortDescriptor.direction === "descending" ? -cmp : cmp;
    });
  }, [sortDescriptor, items]);

  const renderCell = useCallback((order: typeof orders[0], columnKey: React.Key) => {
    const cellValue = order[columnKey as keyof typeof orders[0]];

    switch (columnKey) {
      case "order_id":
        return (
          <div className="flex flex-col">
            <p className="text-bold text-small capitalize">{cellValue.toString()}</p>
            <p className="text-bold text-tiny capitalize text-default-500">{order.order_id}</p>
          </div>
        );
      case "created_at":
        return (
          <div className="flex flex-col">
            <p className="text-bold text-small capitalize">{cellValue.toString()}</p>
            <p className="text-bold text-tiny capitalize text-default-500">{new Date(order.created_at).toDateString()}</p>
          </div>
        );
      case "item_id":
        return (
          <div className="flex flex-col">
            <p className="text-bold text-small capitalize">{cellValue.toString()}</p>
            <p className="text-bold text-tiny capitalize text-default-500">{order.item_id}</p>
          </div>
        );
      case "packing_status":
        return (
          <div className="flex flex-col">
            <p className="text-bold text-small capitalize">{cellValue.toString()}</p>
            <p className="text-bold text-tiny capitalize text-default-500">{order.packing_status}</p>
          </div>
        );
      default:
        return cellValue.toString();
    }
  }, []);

  const onRowsPerPageChange = useCallback((e: React.ChangeEvent<HTMLSelectElement>) => {
    setRowsPerPage(Number(e.target.value));
    setPage(1);
  }, []);

  const onSearchChange = useCallback((value?: string) => {
    if (value) {
      setFilterValue(value);
      setPage(1);
    } else {
      setFilterValue("");
    }
  }, []);

  const topContent = useMemo(() => {
    return (
      <div className="flex flex-col gap-4">
        <div className="flex justify-between gap-3 items-end">
          <Input
            isClearable
            classNames={{
              base: "w-full sm:max-w-[44%]",
              inputWrapper: "border-1",
            }}
            placeholder="Search by order id..."
            size="sm"
            startContent={<SearchIcon className="text-default-300" />}
            value={filterValue}
            variant="bordered"
            onClear={() => setFilterValue("")}
            onValueChange={onSearchChange}
          />
          <div className="flex gap-3">
            <Button
              className="bg-foreground text-background"
              endContent={<PlusIcon />}
              size="sm"
            >
              Add New
            </Button>
          </div>
        </div>
        <div className="flex justify-between items-center">
          <span className="text-default-400 text-small">Total {orders.length} orders</span>
          <label className="flex items-center text-default-400 text-small">
            Rows per page:
            <select
              className="bg-transparent outline-none text-default-400 text-small"
              onChange={onRowsPerPageChange}
            >
              <option value="5">5</option>
              <option value="10">10</option>
              <option value="15">15</option>
            </select>
          </label>
        </div>
      </div>
    );
  }, [filterValue, onSearchChange, onRowsPerPageChange, orders.length]);

  const bottomContent = useMemo(() => {
    return (
      <div className="py-2 px-2 flex justify-between items-center">
        <Pagination
          showControls
          classNames={{
            cursor: "bg-foreground text-background",
          }}
          color="default"
          isDisabled={hasSearchFilter}
          page={page}
          total={pages}
          variant="light"
          onChange={setPage}
        />
        <span className="text-small text-default-400">
          {selectedKeys === "all"
            ? "All items selected"
            : `${selectedKeys.size} of ${items.length} selected`}
        </span>
      </div>
    );
  }, [selectedKeys, items.length, page, pages, hasSearchFilter]);

  const classNames = useMemo(
    () => ({
      wrapper: ["max-h-[382px]", "max-w-3xl"],
      th: ["bg-transparent", "text-default-500", "border-b", "border-divider"],
      td: [
        "group-data-[first=true]:first:before:rounded-none",
        "group-data-[first=true]:last:before:rounded-none",
        "group-data-[middle=true]:before:rounded-none",
        "group-data-[last=true]:first:before:rounded-none",
        "group-data-[last=true]:last:before:rounded-none",
      ],
    }),
    [],
  );

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <Table
      isCompact
      removeWrapper
      aria-label="Example table with custom cells, pagination and sorting"
      bottomContent={bottomContent}
      bottomContentPlacement="outside"
      checkboxesProps={{
        classNames: {
          wrapper: "after:bg-foreground after:text-background text-background",
        },
      }}
      classNames={classNames}
      selectedKeys={selectedKeys}
      selectionMode="multiple"
      sortDescriptor={sortDescriptor}
      topContent={topContent}
      topContentPlacement="outside"
      onSelectionChange={setSelectedKeys}
      onSortChange={setSortDescriptor}
    >
      <TableHeader columns={headerColumns}>
        {(column) => (
          <TableColumn
            key={column.uid}
            align={column.uid === "actions" ? "center" : "start"}
            allowsSorting={column.sortable}
          >
            {column.name}
          </TableColumn>
        )}
      </TableHeader>
      <TableBody emptyContent={"No orders found"} items={sortedItems}>
        {(item) => (
          <TableRow key={item.order_id}>
            {(columnKey) => <TableCell>{renderCell(item, columnKey)}</TableCell>}
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
}

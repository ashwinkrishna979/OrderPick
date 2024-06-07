import ToPackTable from "@/components/tables/to-pack-table";
import { columns } from "@/usecase/data";
import { Table, TableColumn } from "@nextui-org/react";


export default function Home() {


  return (
    <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 w-full">
        <div className="col-span-1">
          <ToPackTable />
        </div>
        <div className="col-span-1">
          <ToPackTable />
        </div>
      </div>
    </section>
  );
}

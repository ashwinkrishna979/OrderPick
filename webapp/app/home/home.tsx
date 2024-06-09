import ToPackTable from "@/components/tables/to-pack-table";
import PackedTable from "@/components/tables/packed-table";


export default function Home() {


  return (
    <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 w-full">
        <div className="col-span-2">
        <br/>
          <h2>Orders To Pack</h2>
          <br/>
          <ToPackTable />
        </div>
        <div className="col-span-2">
        <br/>
        <h2>Packed Orders</h2>
        <br/>
          <PackedTable />
        </div>
      </div>
    </section>
  );
}

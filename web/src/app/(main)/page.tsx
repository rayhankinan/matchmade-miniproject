import DiscoverMovie from "~/components/app/main/home/discover-page";
import SearchMovie from "~/components/app/main/home/search-page";

export default function HomePage({
  searchParams,
}: {
  searchParams: {
    q?: string;
  };
}) {
  const { q: query } = searchParams;

  return (
    <main className="flex min-h-screen flex-col items-center justify-center gap-12">
      {query ? <SearchMovie query={query} /> : <DiscoverMovie />}
    </main>
  );
}

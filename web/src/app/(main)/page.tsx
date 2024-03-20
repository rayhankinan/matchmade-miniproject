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
      {query ? (
        <h1 className="pt-20 text-2xl font-bold">
          Search results for &quot;{query}&quot;
        </h1>
      ) : (
        <h1 className="pt-20 text-5xl font-bold">Discover Movies</h1>
      )}
      <div className="flex min-h-screen flex-col items-center justify-center">
        {query ? <SearchMovie query={query} /> : <DiscoverMovie />}
      </div>
    </main>
  );
}

import DiscoverWatchlist from "~/components/app/main/watchlist/discover-watchlist-page";
import SearchWatchlist from "~/components/app/main/watchlist/search-watchlist-page";
import WatchlistAlert from "~/components/app/warning/watchlist-alert";
import getSession from "~/server/auth";

export default function WatchlistPage({
  searchParams,
}: {
  searchParams: {
    q?: string;
  };
}) {
  const { q: title } = searchParams;

  const session = getSession();

  return (
    <main className="flex min-h-screen flex-col items-center justify-center gap-12">
      {session &&
        (title ? (
          <h1 className="pt-20 text-2xl font-bold">
            Search results for{" "}
            <span className="italic">&quot;{title}&quot;</span>
          </h1>
        ) : (
          <h1 className="pt-20 text-2xl font-bold">Discover Watchlist</h1>
        ))}
      {session ? (
        <div className="flex min-h-screen flex-col items-center justify-center">
          {title ? <SearchWatchlist title={title} /> : <DiscoverWatchlist />}
        </div>
      ) : (
        <WatchlistAlert />
      )}
    </main>
  );
}

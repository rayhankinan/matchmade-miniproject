import Profile from "~/components/app/navbar/profile";
import ModeToggle from "~/components/app/navbar/mode-toggle";
import Link from "next/link";
import Image from "next/image";

export default function AppNavbar({
  logoSrc,
  logoBase64,
  searchBar,
  displayProfile = false,
}: {
  logoSrc: string;
  logoBase64: string;
  searchBar?: React.ReactElement;
  displayProfile?: boolean;
}) {
  return (
    <header className="fixed top-0 z-10 flex h-20 w-full shrink-0 items-center justify-between px-4 md:px-6">
      <div className="flex flex-row items-center gap-2">
        <Image
          src={logoSrc}
          alt="The Movie Watchlist logo"
          width={40}
          height={40}
          placeholder="blur"
          blurDataURL={logoBase64}
        />
        <Link href="/">
          <span className="cursor-pointer text-2xl font-bold">
            The Movie Watchlist
          </span>
        </Link>
      </div>
      <div className="flex flex-row items-center gap-2">
        {searchBar}
        <ModeToggle />
        {displayProfile && <Profile />}
      </div>
    </header>
  );
}

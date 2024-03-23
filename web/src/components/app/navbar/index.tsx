import Image from "next/image";

import Profile from "~/components/app/navbar/profile";
import ModeToggle from "~/components/app/navbar/mode-toggle";
import AppNavigationMenu from "~/components/app/navbar/navigation-menu";
import SearchBar from "~/components/app/navbar/search-bar";

export default function AppNavbar({
  logoSrc,
  logoBase64,
  displayNavigationMenu = false,
  displaySearchBar = false,
  displayProfile = false,
}: {
  logoSrc: string;
  logoBase64: string;
  displayNavigationMenu?: boolean;
  displaySearchBar?: boolean;
  displayProfile?: boolean;
}) {
  return (
    <header className="fixed top-0 z-10 flex h-20 w-full shrink-0 items-center justify-between px-4 md:px-6">
      <div className="flex flex-row items-center gap-4">
        <div className="flex flex-row items-center gap-2">
          <Image
            src={logoSrc}
            alt="The Movie Watchlist logo"
            width={60}
            height={60}
            placeholder="blur"
            blurDataURL={logoBase64}
          />
        </div>
        {displayNavigationMenu && <AppNavigationMenu />}
      </div>
      <div className="flex flex-row items-center gap-2">
        {displaySearchBar && <SearchBar />}
        <ModeToggle />
        {displayProfile && <Profile />}
      </div>
    </header>
  );
}

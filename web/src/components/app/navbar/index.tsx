import Profile from "~/components/app/navbar/profile";
import ModeToggle from "~/components/app/navbar/mode-toggle";

export default function AppNavbar({
  searchBar,
  displayProfile = false,
}: {
  searchBar?: React.ReactElement;
  displayProfile?: boolean;
}) {
  return (
    <header className="fixed top-0 z-10 flex justify-between h-20 w-full shrink-0 items-center px-4 md:px-6">
        <div className="flex">
          {searchBar}
        </div>
        <div className="flex gap-2">
          <ModeToggle />
          {displayProfile && <Profile />}
        </div>
    </header>
  );
}

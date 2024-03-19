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
    <header className="sticky flex h-20 w-full shrink-0 items-center px-4 md:px-6">
      <div className="ml-auto flex gap-2">
        {displayProfile && <Profile />}
        <ModeToggle />
        {searchBar}
      </div>
    </header>
  );
}

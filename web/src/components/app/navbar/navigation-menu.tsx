import Link from "next/link";

import {
  navigationMenuTriggerStyle,
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
} from "~/components/ui/navigation-menu";

import useSession from "~/hooks/auth";

export default function AppNavigationMenu() {
  const session = useSession();

  return (
    <NavigationMenu>
      <NavigationMenuList>
        <NavigationMenuItem>
          <Link href="/" legacyBehavior passHref>
            <NavigationMenuLink className={navigationMenuTriggerStyle()}>
              <span className="text-lg font-bold">Home</span>
            </NavigationMenuLink>
          </Link>
        </NavigationMenuItem>

        {session && (
          <NavigationMenuItem>
            <Link href="/watchlist" legacyBehavior passHref>
              <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                <span className="text-lg font-bold">Watchlist</span>
              </NavigationMenuLink>
            </Link>
          </NavigationMenuItem>
        )}
      </NavigationMenuList>
    </NavigationMenu>
  );
}

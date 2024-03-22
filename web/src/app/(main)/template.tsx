import MovieTemplate from "~/components/app/template/MovieTemplate";
import getBlurData from "~/lib/get-blur-data";

export default async function RootTemplate({
  children,
}: React.PropsWithChildren) {
  const src = "/images/raster/logo.png";

  const { base64 } = await getBlurData("./public/images/raster/logo.png", {
    size: 16,
  });

  return (
    <MovieTemplate logoSrc={src} logoBase64={base64}>
      {children}
    </MovieTemplate>
  );
}

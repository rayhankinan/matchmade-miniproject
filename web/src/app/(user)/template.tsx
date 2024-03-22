import AuthTemplate from "~/components/app/template/AuthTemplate";
import getBlurData from "~/lib/get-blur-data";

export default async function RootTemplate({
  children,
}: React.PropsWithChildren) {
  const src = "/images/raster/logo.png";

  const { base64 } = await getBlurData("./public/images/raster/logo.png", {
    size: 16,
  });

  return (
    <AuthTemplate logoSrc={src} logoBase64={base64}>
      {children}
    </AuthTemplate>
  );
}

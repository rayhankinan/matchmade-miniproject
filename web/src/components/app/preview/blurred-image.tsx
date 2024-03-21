import Image from "next/image";

import getBlurData from "~/lib/get-blur-data";

export default async function BlurredImage({
  src,
  alt,
}: {
  src: string;
  alt: string;
}) {
  const { base64 } = await getBlurData(src);

  return (
    <Image
      src={src}
      alt={alt}
      placeholder="blur"
      blurDataURL={base64}
      sizes="(min-width: 500px) 50vw, 100vw"
      fill
      priority
      className="rounded-sm"
    />
  );
}

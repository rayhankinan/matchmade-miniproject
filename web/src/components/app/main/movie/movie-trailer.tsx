"use client";

import { useEffect, useState, useRef } from "react";

import Spinner from "~/components/app/icon/spinner";
import { env } from "~/env";

export default function MovieTrailer({
  data: { name, site, key },
}: {
  data: {
    name: string;
    site: "YouTube" | "Vimeo";
    key: string;
  };
}) {
  const [load, setLoad] = useState(false);

  const observerTarget = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const currentTarget = observerTarget.current;

    const observer = new IntersectionObserver((entries) => {
      if (entries[0] && entries[0].isIntersecting) {
        setLoad(true);
        observer.disconnect();
      }
    });

    if (currentTarget) observer.observe(currentTarget);

    return () => {
      if (currentTarget) observer.unobserve(currentTarget);
    };
  }, []);

  return (
    <div ref={observerTarget} className="container mx-auto size-fit">
      {load ? (
        <iframe
          src={`${site === "YouTube" ? env.NEXT_PUBLIC_YOUTUBE_VIDEO_URL : env.NEXT_PUBLIC_VIMEO_VIDEO_URL}/${key}`}
          title={name}
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
          allowFullScreen
        ></iframe>
      ) : (
        <Spinner />
      )}
    </div>
  );
}

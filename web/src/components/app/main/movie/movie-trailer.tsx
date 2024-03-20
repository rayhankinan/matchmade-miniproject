"use client";

import { useEffect, useState, useRef } from "react";
import ReactPlayer from "react-player";

import Spinner from "~/components/app/icon/spinner";
import { env } from "~/env";

export default function MovieTrailer({
  data: { site, key },
}: {
  data: {
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
    <div ref={observerTarget}>
      {load ? (
        <ReactPlayer
          className="w-full max-w-md"
          url={
            site === "YouTube"
              ? `${env.NEXT_PUBLIC_YOUTUBE_VIDEO_URL}?v=${key}`
              : `${env.NEXT_PUBLIC_VIMEO_VIDEO_URL}/${key}`
          }
        />
      ) : (
        <Spinner />
      )}
    </div>
  );
}

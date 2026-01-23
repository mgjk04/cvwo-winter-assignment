"use client";
import { use } from "react";
import useReadPost from "../../_hooks/useReadPost";
import EditPost from "./_components/editPost";

export default function Page({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const readURL = `${process.env.NEXT_PUBLIC_API_URL}/posts/${id}`;
  const { data, status } = useReadPost(readURL);
  if (status === "error") {
    return <>this is an error page</>; //TODO ADD ERROR PAGE
  }
  return <EditPost submitURL={readURL} post={data} />;
}

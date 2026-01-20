"use client";
import { use } from "react";
import ModifyComment from "./_components/modifyComment";
import useReadComment from "../../_hooks/useReadComment";

export default function Page({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const readURL = `${process.env.NEXT_PUBLIC_API_URL}/comments/${id}`;
  const { data, status } = useReadComment(readURL);
  if (status === "error") {
    return <>this is an error page</>; //TODO ADD ERROR PAGE
  }
  return <ModifyComment submitURL={readURL} comment={data} />;
}

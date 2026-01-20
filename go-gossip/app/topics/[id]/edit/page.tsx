"use client";
import { use } from "react";
import useReadTopic from "../../_hooks/useReadTopic";
import EditTopic from "./_components/editTopic";

export default function Page({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const readURL = `${process.env.NEXT_PUBLIC_API_URL}/topics/${id}`;
  const editURL = readURL;
  const { data, status } = useReadTopic(readURL);
  if (status === "error") {
    return <>this is an error page</>; //TODO ADD ERROR PAGE
  }
  return <EditTopic submitURL={editURL} topic={data} />;
}

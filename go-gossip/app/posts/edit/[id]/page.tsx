"use client"
import { use } from "react";
import useReadPost from "../../_hooks/useReadPost";
import { verb } from "../../types";
import ModifyPost from "../../_components/modifyPost";

export default function Page({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const readURL = "http://localhost:8080/posts/" + id;
  const { data, status } = useReadPost(readURL);
  if (status === "error") {
    return <>this is an error page</>; //TODO ADD ERROR PAGE
  }
  return <ModifyPost verb={verb.Edit} submitURL={readURL} topic={data} />;
}

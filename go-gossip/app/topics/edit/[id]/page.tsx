"use client";
import { verb } from "../../types";
import { use } from "react";
import useReadTopic from "../../_hooks/useReadTopic";
import ModifyTopic from "../../_components/modifyTopic";

export default function Page({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const readURL = "http://localhost:8080/topics/" + id;
  const { data, status } = useReadTopic(readURL);
  if (status === "error") {
    return <>this is an error page</>; //TODO ADD ERROR PAGE
  }
  return <ModifyTopic verb={verb.Edit} submitURL={readURL} topic={data} />;
}

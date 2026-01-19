"use client";
import ModifyTopic from "../_components/modifyTopic";
import { verb } from "../types";

export default function Page() {
  const createTopicURL = "http://localhost:8080/topics/";
  return <ModifyTopic verb={verb.Create} submitURL={createTopicURL} />;
}

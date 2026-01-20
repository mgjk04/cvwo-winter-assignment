"use client";
import CreateTopic from "./_components/createTopic";
import "dotenv/config";

export default function Page() {
  const createTopicURL = process.env.API_URL + "/topics/";
  return <CreateTopic submitURL={createTopicURL} />;
}

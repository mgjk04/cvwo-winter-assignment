"use client";
import {
    Box,
  Link,
  List,
  ListItemButton,
  ListItemText,
  Pagination,
  Stack,
  Typography,
} from "@mui/material";
import { useState } from "react";
import useReadTopic from "./_hooks/useReadTopic";
import { dataState, topic } from "./types";

export default function TopicsPage() {
  const [dataState, setDataState] = useState<dataState>({
      page: 1,
      limit: 10,
  });
  const readURL =
    process.env.NEXT_PUBLIC_API_URL +
    `/topics/?page=${dataState.page}&limit=${dataState.limit}`;
  const { data, isPending, isError } = useReadTopic(readURL);
  return (
    <Stack>
        <List>
        {(data || { topics: [], count: 0 }).topics.map((t: topic) => {
            const topicURL = `/topics/${t.id}`;
            return (
            <ListItemButton key={t.id} LinkComponent={Link} href={topicURL}>
                <ListItemText>
                <Typography variant="h5">{t.topicname}</Typography>
                </ListItemText>
            </ListItemButton>
            );
        })}
        </List>
    </Stack>
  );
}

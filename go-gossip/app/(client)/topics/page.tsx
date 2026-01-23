"use client";
import {
  Button,
  Grid,
  Link,
  List,
  ListItemButton,
  ListItemText,
  Stack,
  Typography,
} from "@mui/material";
import { useEffect, useState } from "react";
import useReadTopic from "./_hooks/useReadTopic";
import { dataState, topic } from "./types";
import { getCookie } from "cookies-next";
import NavigateNextIcon from "@mui/icons-material/NavigateNext";
import NavigateBeforeIcon from "@mui/icons-material/NavigateBefore";

export default function TopicsPage() {
  const [dataState, setDataState] = useState<dataState>({
    page: 1,
    limit: 10,
  });
  const [userState, setUserState] = useState<userState>({
    username: "",
    userId: "",
  });
  useEffect(() => {
    const username = getCookie("username") || "";
    const userId = getCookie("user_id") || "";
    // eslint-disable-next-line react-hooks/set-state-in-effect
    setUserState({
      username: username as string,
      userId: userId as string,
    });
  }, []);
  const readURL = `${process.env.NEXT_PUBLIC_API_URL}/topics/?page=${dataState.page}&limit=${dataState.limit}`;
  const createURL = `/topics/create`;
  const isSignedIn = userState.userId != "";
  const { data, isPending, isError, refetch } = useReadTopic(readURL);
  const info = data || { topics: [], count: 0 };
  return (
    <Stack>
      <Grid>
        {isSignedIn && (
          <Button component={Link} href={createURL}>
            Start a Topic!
          </Button>
        )}
      </Grid>
      {info.count === 0 ? (
        <Grid container justifyContent="center">
          <Typography variant="body2">Wow such empty...</Typography>
        </Grid>
      ) : (
        <List>
          {(data || { topics: [], count: 0 }).topics.map((t: topic) => {
            const topicURL = `/topics/${t.id}`;
            return (
              <ListItemButton key={t.id} LinkComponent={Link} href={topicURL}>
                <ListItemText>
                  <Typography variant="h5">{t.topicname}</Typography>
                  <Typography variant="body2">{t.description}</Typography>
                </ListItemText>
              </ListItemButton>
            );
          })}
        </List>
      )}
      <Grid container justifyContent="center">
        <Button
          disabled={dataState.page === 1}
          onClick={() => {
            setDataState({ ...dataState, page: dataState.page - 1 });
          }}
        >
          <NavigateBeforeIcon />
        </Button>
        <Button
          disabled={info.count === 0}
          onClick={() => {
            setDataState({ ...dataState, page: dataState.page + 1 });
          }}
        >
          <NavigateNextIcon />
        </Button>
      </Grid>
    </Stack>
  );
}

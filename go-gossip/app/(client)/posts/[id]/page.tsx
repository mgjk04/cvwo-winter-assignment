"use client";
import {
  Button,
  Card,
  CardContent,
  CardHeader,
  Chip,
  Grid,
  Link,
  List,
  ListItemText,
  Stack,
  Typography,
} from "@mui/material";
import { use, useEffect, useState } from "react";
import useReadPost from "../_hooks/useReadPost";
import useReadComment from "@/app/(client)/comments/_hooks/useReadComment";
import { comment } from "@/app/(client)/comments/types";
import dayjs from "dayjs";
import { getCookie } from "cookies-next";
import { userState } from "@/app/types";
import CreateComment from "./create/_components/createComment";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";
import useDeleteComment from "../../comments/_hooks/useDeleteComment";
import useDeletePost from "../_hooks/useDeletePost";
import { useRouter } from "next/navigation";
import NavigateNextIcon from "@mui/icons-material/NavigateNext"
import NavigateBeforeIcon from "@mui/icons-material/NavigateBefore"

export default function PostPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = use(params);
  const [dataState, setDataState] = useState<dataState>({
    page: 1,
    limit: 10,
  });
  const [userState, setUserState] = useState<userState>({
    username: "",
    userId: "",
  });
  const router = useRouter();

  useEffect(() => {
    const username = getCookie("username") || "";
    const userId = getCookie("user_id") || "";
    // eslint-disable-next-line react-hooks/set-state-in-effect
    setUserState({
      username: username as string,
      userId: userId as string,
    });
  }, []);

  const readPostURL = `${process.env.NEXT_PUBLIC_API_URL}/posts/${id}`;
  const readCommentURL = `${process.env.NEXT_PUBLIC_API_URL}/posts/${id}/comments?page=${dataState.page}&limit=${dataState.limit}`;
  const createCommentURL = `${process.env.NEXT_PUBLIC_API_URL}/posts/${id}/comments`;
  const editPostURL = `/posts/${id}/edit`;
  const deletePostURL = readPostURL;
  const {
    data: postData,
    isPending: postIsPending,
    isError: postIsError,
  } = useReadPost(readPostURL);
  const {
    data: commentData,
    isPending: commentIsPending,
    isError: commentIsError,
  } = useReadComment(readCommentURL);
  const { mutate: deleteComment } = useDeleteComment();
  const { mutate: deletePost } = useDeletePost(deletePostURL, router);
  const info = commentData || { comments: [], count: 0 };

  return (
    <Stack>
      <Card>
        <CardContent>
          <Grid size="grow" justifyContent="space-between" container>
            <Grid>
              <Stack spacing={1}>
                <Typography className="wrap-break-word" variant="h5">
                  {postData?.title || ""}
                </Typography>
                <Typography className="wrap-break-word" variant="body1">
                  {postData?.description || ""}
                </Typography>
                <Grid container spacing={1}>
                  <Chip
                    className="max-w-max"
                    label={`By ${postData?.authorname}`}
                    size="small"
                  />
                  <Chip
                    className="max-w-max"
                    label={`On ${dayjs(postData?.created_at).format("HH:mm DD-MM-YYYY")}`}
                    size="small"
                  />
                </Grid>
              </Stack>
            </Grid>
            <Grid size="auto">
              {userState?.userId &&
                userState.userId === postData?.author_id && (
                  <Grid>
                    <Button component={Link} href={editPostURL}>
                      <EditIcon />
                    </Button>
                    <Button
                      onClick={() =>
                        deletePost(undefined, {
                          onSuccess: () => router.back(),
                        })
                      }
                    >
                      <DeleteIcon />
                    </Button>
                  </Grid>
                )}
            </Grid>
          </Grid>
          {userState?.userId && (
            <CreateComment submitURL={createCommentURL} />
          )}
        </CardContent>
      </Card>

      <List>
        {info.count === 0 ? (
          <Typography variant="body2">Wow such empty...</Typography>
        ) : (
          info.comments?.map((c: comment) => {
            const editCommentURL = `/comments/${c.id}/edit`;
            return (
              <ListItemText key={c.id}>
                <Card>
                  <CardHeader
                    title={c.authorname}
                    action={
                      userState?.userId &&
                      userState.userId === c?.author_id && (
                        <Grid>
                          <Button
                            size="small"
                            component={Link}
                            href={editCommentURL}
                          >
                            <EditIcon />
                          </Button>
                          <Button
                            onClick={() => {
                              const deleteCommentURL = `${process.env.NEXT_PUBLIC_API_URL}/comments/${c.id}`;
                              deleteComment(deleteCommentURL);
                            }}
                          >
                            <DeleteIcon />
                          </Button>
                        </Grid>
                      )
                    }
                  />

                  <CardContent>
                    <Typography className="wrap-break-word" variant="body1">
                      {c.content}
                    </Typography>
                    <Chip
                      label={dayjs(c.created_at)
                        .format("HH:mm DD-MM-YYYY")
                        .toString()}
                    />
                  </CardContent>
                </Card>
              </ListItemText>
            );
          })
        )}
      </List>
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

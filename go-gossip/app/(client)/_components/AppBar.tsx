"use client";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Button from "@mui/material/Button";
import { Grid, Link } from "@mui/material";
import { deleteCookie, getCookie } from "cookies-next";
import useLogout from "../(auth)/_hooks/useLogout";
import { MouseEvent, useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import { usePathname } from "next/navigation";

export default function ButtonAppBar() {
  const [userState, setUserState] = useState<userState>({
    username: "",
    userId: "",
  });
  const [historySize, setHistorySize] = useState(0);
  useEffect(() => {
    const username = getCookie("username") || "";
    const userId = getCookie("user_id") || "";
    // eslint-disable-next-line react-hooks/set-state-in-effect
    setUserState({
      username: username as string,
      userId: userId as string,
    });
    setHistorySize(window.history.length);
  }, []);
  const pathName = usePathname();
  const router = useRouter();
  const { mutate } = useLogout();
  const isLoginOrSignupPage = pathName == "/login" || pathName == "/signup";
  const showBackArrowCond = historySize > 0 && pathName != "/topics";
  const showLoginButtonCond =
    !userState.userId && !userState.username && !isLoginOrSignupPage;
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          {showBackArrowCond && (
            <Button variant="contained" color="secondary" onClick={(event: MouseEvent) => router.back()}>
              <ArrowBackIcon />
            </Button>
          )}
          {!isLoginOrSignupPage ? (
            showLoginButtonCond ? (
              <Grid container spacing={1}>
                <Button variant="contained" color="secondary" onClick={() => router.push("/login")}>Login</Button>
                <Button variant="contained" color="secondary" onClick={() => router.push("/signup")}>Sign up</Button>
              </Grid>
            ) : (
              <Button
                onClick={(event: MouseEvent) => {
                  mutate(undefined, {
                    onSettled: () => {
                      deleteCookie("user_id");
                      deleteCookie("username");
                      deleteCookie("access_token"); //just in case API fails
                      deleteCookie("refresh_token");
                      setUserState({ userId: "", username: "" });
                    },
                  });
                }}
              >
                Logout
              </Button>
            )
          ) : (
            <></>
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
}

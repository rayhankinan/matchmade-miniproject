interface JwtPayload {
  userID: string;
  email: string;
  username: string;
  exp: number;
}

export default JwtPayload;

CREATE TABLE [users] (
    [ID]   bigint PRIMARY KEY,
    [ACID] character varying(max),
    [TGID] bigint,
    [BID]  bigint FOREIGN KEY REFERENCES [battery](ID),
    [CID]  bigint FOREIGN KEY REFERENCES [clicks](ID),
    [SID]  bigint FOREIGN KEY REFERENCES [score],
)
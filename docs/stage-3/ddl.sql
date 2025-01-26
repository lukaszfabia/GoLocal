CREATE TABLE User3 ("Column" SERIAL NOT NULL, PRIMARY KEY ("Column"));

CREATE TABLE Location (
    ID SERIAL NOT NULL,
    City varchar(255),
    Country varchar(255),
    Zip varchar(255),
    AddressID int4 NOT NULL,
    CoordsID int4 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE Address (
    ID SERIAL NOT NULL,
    Street varchar(255),
    StreetNumber varchar(255),
    Info varchar(255),
    PRIMARY KEY (ID)
);

CREATE TABLE Coords (
    ID SERIAL NOT NULL,
    Latitude float4 NOT NULL,
    Longitude float4 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE Event (
    ID SERIAL NOT NULL,
    Title varchar(255),
    "Desc" varchar(2550),
    ImageURL varchar(255),
    StartDate int4,
    FinishDate int4,
    IsAdultOnly bool NOT NULL,
    EventType varchar(255),
    PRIMARY KEY (ID)
);

CREATE TABLE "Tag" (
    ID SERIAL NOT NULL,
    Name varchar(255),
    PRIMARY KEY (ID)
);

CREATE TABLE PreferenceSurveyQuestion (
    ID SERIAL NOT NULL,
    Text varchar(255),
    QuestionType varchar(32),
    PreferenceSurveyID int4 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE "User" (
    ID SERIAL NOT NULL,
    FirstName varchar(255),
    LastName varchar(255),
    Email varchar(255),
    Password varchar(255),
    Birthday int4,
    Bio varchar(255),
    AuthProvider varchar(255),
    IsVerified bool NOT NULL,
    IsPremium bool NOT NULL,
    AvatarURL varchar(255),
    PRIMARY KEY (ID)
);

CREATE TABLE Opinion (
    ID SERIAL NOT NULL,
    Text varchar(255),
    UserID int4 NOT NULL,
    EventID int4 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE PreferenceSurveyAnswer (
    ID SERIAL NOT NULL,
    PreferenceSurveyID int4 NOT NULL,
    UserID int4 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE Vote (
    ID SERIAL NOT NULL,
    Text varchar(255),
    EndDate time,
    EventID int4 NOT NULL,
    UserID int4 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE PreferenceSurvey (
    ID SERIAL NOT NULL,
    Title varchar(255),
    Description varchar(255),
    PRIMARY KEY (ID)
);

CREATE TABLE VoteOption (
    ID SERIAL NOT NULL,
    Text varchar(255),
    VoteID int4 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE VoteAnswer (
    ID SERIAL NOT NULL,
    UserID int4 NOT NULL,
    VoteID int4 NOT NULL,
    VoteOptionID int4 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE Comment (
    ID SERIAL NOT NULL,
    Content varchar(2550),
    UserID int4 NOT NULL,
    EventID int4 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE PreferenceSurveyOption (
    ID SERIAL NOT NULL,
    PreferenceSurveyQuestionID int4 NOT NULL,
    Text varchar(255),
    TagID int4 NOT NULL,
    TagPositive bool NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE PreferenceSurveyAnswerOption (
    ID SERIAL NOT NULL,
    "PreferenceSurveyAnswer ID" int4 NOT NULL,
    "PreferenceSurveyOption ID" int4 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE Vote2 (
    ID2 SERIAL NOT NULL,
    Attribute int4 NOT NULL,
    ID int4 NOT NULL,
    Question varchar(255),
    EventID int4 NOT NULL,
    Type int4,
    EndDate int4,
    PRIMARY KEY (ID2)
);

CREATE TABLE VoteOptione (
    ID2 SERIAL NOT NULL,
    Attribute int4 NOT NULL,
    ID int4 NOT NULL,
    VoteID int4 NOT NULL,
    Text varchar(255),
    OptionID int4,
    PRIMARY KEY (ID2)
);

CREATE TABLE VoteAnswer2 (
    ID2 SERIAL NOT NULL,
    Attribute int4 NOT NULL,
    ID int4 NOT NULL,
    VoteID int4 NOT NULL,
    UserID int4 NOT NULL,
    VoteOptionID int4 NOT NULL,
    PRIMARY KEY (ID2)
);

CREATE TABLE Location_Event (
    LocationID int4 NOT NULL,
    EventID int4 NOT NULL,
    PRIMARY KEY (LocationID, EventID)
);

CREATE TABLE Event_Organizer (
    EventID int4 NOT NULL,
    UserID int4 NOT NULL,
    PRIMARY KEY (EventID, UserID)
);

CREATE TABLE Tag_Event (
    TagID int4 NOT NULL,
    EventID int4 NOT NULL,
    PRIMARY KEY (TagID, EventID)
);

CREATE TABLE Preference (
    TagID int4 NOT NULL,
    UserID int4 NOT NULL,
    PRIMARY KEY (TagID, UserID)
);

ALTER TABLE
    "PreferenceSurveyOption "
ADD
    CONSTRAINT FKPreference697419 FOREIGN KEY (PreferenceSurveyQuestionID) REFERENCES PreferenceSurveyQuestion (ID);

ALTER TABLE
    "PreferenceSurveyAnswerOption "
ADD
    CONSTRAINT FKPreference185068 FOREIGN KEY ("PreferenceSurveyAnswer ID") REFERENCES "PreferenceSurveyAnswer " (ID);

ALTER TABLE
    Location_Event
ADD
    CONSTRAINT FKLocation_E773212 FOREIGN KEY (LocationID) REFERENCES Location (ID);

ALTER TABLE
    Location_Event
ADD
    CONSTRAINT FKLocation_E759979 FOREIGN KEY (EventID) REFERENCES Event (ID);

ALTER TABLE
    "PreferenceSurveyAnswer "
ADD
    CONSTRAINT FKPreference401108 FOREIGN KEY (PreferenceSurveyID) REFERENCES PreferenceSurvey (ID);

ALTER TABLE
    Vote
ADD
    CONSTRAINT FKVote164762 FOREIGN KEY (EventID) REFERENCES Event (ID);

ALTER TABLE
    Event_Organizer
ADD
    CONSTRAINT FKEvent_Orga238042 FOREIGN KEY (EventID) REFERENCES Event (ID);

ALTER TABLE
    Event_Organizer
ADD
    CONSTRAINT FKEvent_Orga723989 FOREIGN KEY (UserID) REFERENCES "User" (ID);

ALTER TABLE
    VoteOption
ADD
    CONSTRAINT FKVoteOption399579 FOREIGN KEY (VoteID) REFERENCES Vote (ID);

ALTER TABLE
    Tag_Event
ADD
    CONSTRAINT FKTag_Event81923 FOREIGN KEY (TagID) REFERENCES "Tag" (ID);

ALTER TABLE
    Tag_Event
ADD
    CONSTRAINT FKTag_Event757391 FOREIGN KEY (EventID) REFERENCES Event (ID);

ALTER TABLE
    Location
ADD
    CONSTRAINT FKLocation196189 FOREIGN KEY (AddressID) REFERENCES Address (ID);

ALTER TABLE
    Location
ADD
    CONSTRAINT FKLocation116422 FOREIGN KEY (CoordsID) REFERENCES Coords (ID);

ALTER TABLE
    VoteAnswer
ADD
    CONSTRAINT FKVoteAnswer331514 FOREIGN KEY (UserID) REFERENCES "User" (ID);

ALTER TABLE
    VoteAnswer
ADD
    CONSTRAINT FKVoteAnswer727377 FOREIGN KEY (VoteID) REFERENCES Vote (ID);

ALTER TABLE
    VoteAnswer
ADD
    CONSTRAINT FKVoteAnswer906721 FOREIGN KEY (VoteOptionID) REFERENCES VoteOption (ID);

ALTER TABLE
    Vote
ADD
    CONSTRAINT FKVote844796 FOREIGN KEY (UserID) REFERENCES "User" (ID);

ALTER TABLE
    Opinion
ADD
    CONSTRAINT FKOpinion353938 FOREIGN KEY (UserID) REFERENCES "User" (ID);

ALTER TABLE
    Opinion
ADD
    CONSTRAINT FKOpinion655620 FOREIGN KEY (EventID) REFERENCES Event (ID);

ALTER TABLE
    Comment
ADD
    CONSTRAINT FKComment537260 FOREIGN KEY (UserID) REFERENCES "User" (ID);

ALTER TABLE
    Comment
ADD
    CONSTRAINT FKComment424771 FOREIGN KEY (EventID) REFERENCES Event (ID);

ALTER TABLE
    "PreferenceSurveyAnswer "
ADD
    CONSTRAINT FKPreference909423 FOREIGN KEY (UserID) REFERENCES "User" (ID);

ALTER TABLE
    PreferenceSurveyQuestion
ADD
    CONSTRAINT FKPreference664694 FOREIGN KEY (PreferenceSurveyID) REFERENCES PreferenceSurvey (ID);

ALTER TABLE
    "PreferenceSurveyAnswerOption "
ADD
    CONSTRAINT FKPreference355156 FOREIGN KEY ("PreferenceSurveyOption ID") REFERENCES "PreferenceSurveyOption " (ID);

ALTER TABLE
    Preference
ADD
    CONSTRAINT FKPreference286859 FOREIGN KEY (TagID) REFERENCES "Tag" (ID);

ALTER TABLE
    Preference
ADD
    CONSTRAINT FKPreference883384 FOREIGN KEY (UserID) REFERENCES "User" (ID);
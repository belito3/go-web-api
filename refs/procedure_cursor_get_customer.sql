--------------------------------------------------------
--  File created - Friday-July-31-2020   
--------------------------------------------------------
--------------------------------------------------------
--  DDL for Procedure GET_CUSTOMER
--------------------------------------------------------
set define off;

  CREATE OR REPLACE EDITIONABLE PROCEDURE "GET_CUSTOMER" (
    pID             IN NUMBER,
    oTOTAL_ELEMENT  OUT NUMBER,
    oCUR            OUT SYS_REFCURSOR
)
AS 
V_COUNT NUMBER := 0;

BEGIN
    OPEN oCUR FOR
--    SELECT ID as "id", NAME as "name", AGE as "age", CREATED as "created" FROM test WHERE id=pID OFFSET 0 ROWS FETCH NEXT 1 ROWS ONLY;
    SELECT ID as "id", NAME as "name", AGE as "age", CREATED as "created" FROM test OFFSET 0 ROWS FETCH NEXT 3 ROWS ONLY;    
    SELECT COUNT(ID) INTO V_COUNT FROM test;
    oTOTAL_ELEMENT := V_COUNT;
    
--    SELECT oTOTAL_ELEMENT, oCUR FROM dual;
END GET_CUSTOMER;

/

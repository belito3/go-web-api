--------------------------------------------------------
--  File created - Saturday-May-16-2020   
--------------------------------------------------------
--------------------------------------------------------
--  DDL for Procedure UPSERT_TEST
--------------------------------------------------------
set define off;

  CREATE OR REPLACE EDITIONABLE PROCEDURE "UPSERT_TEST" (
   p_id IN test.id%TYPE,
   p_name IN test.name%TYPE,
   p_age IN test.age%TYPE,
   p_created IN test.created%TYPE)
AS 
BEGIN
  INSERT INTO test (ID, NAME, AGE, CREATED) VALUES (p_id, p_name, p_age, p_created);
  EXCEPTION 
  WHEN DUP_VAL_ON_INDEX THEN 
  UPDATE test SET NAME = p_name, AGE = p_age WHERE ID = p_id;
END UPSERT_TEST;

/

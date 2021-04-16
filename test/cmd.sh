grpcurl -d '{"spids": [28447878], "model": {"simple": {}}, "params": {"version": "tikinow_v1_20210407"}, "user_metadata": {"trackity_id": "cc46d901-f489-ea8e-5890-1d5bfd365526"}}' -plaintext localhost:9091  tiki.smarter.servings.v1.services.GeneralComputationService/Rank

SELECT * FROMrom p_reco_tikinow_v1_20210407 WHERE key="9e776a76-8848-4b09-99de-d8bd3ba33958";
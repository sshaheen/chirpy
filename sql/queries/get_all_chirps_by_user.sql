-- name: GetAllChirpsByUser :many
SELECT *
FROM chirps
WHERE user_id = $1
ORDER BY 
  CASE 
    WHEN $2 = 'desc' THEN NULL
    ELSE created_at
  END ASC,
  CASE 
    WHEN $2 = 'desc' THEN created_at
    ELSE NULL
  END DESC;
aws ec2 describe-images --image-ids | jq -r '.Images | .[] | select(.OwnerId=="999999999") | select(.Public==false) | .ImageId' | xargs -t -P 8 -I {} sh -c 'aws ec2 deregister-image --image-id {} || echo Deregistering {} failed'
aws ec2 describe-images --image-ids --region us-east-1 | jq -r '.Images | .[] | select(.OwnerId=="999999999") | select(.Public==false) | .ImageId' | xargs -t -P 8 -I {} sh -c 'aws ec2 deregister-image --image-id {} || echo Deregistering {} failed'

aws ec2 describe-snapshots | jq -r '.Snapshots | .[] | select(.OwnerId="999999999") | .SnapshotId' | xargs -t -I {} sh -c 'aws ec2 delete-snapshot --snapshot-id {} || echo Deleting snapshot {} failed'
